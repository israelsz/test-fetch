import csv
import pymongo
from pymongo import MongoClient
from bson import ObjectId
import datetime

# Conexión a la base de datos MongoDB
client = MongoClient('mongodb://localhost:27017/')
db = client["fusupo"]

# seleccionar colección
competencia = db['Competencia']
formularioCompetencia = db['FormularioCompetencia']
cargo = db['Cargo']
equipo = db['Equipo']
usuario = db['User']

def populate_competencia(nombre):
    n=0
    with open(nombre, newline='') as csvfile:
        
        reader = csv.reader(csvfile, delimiter=',', quotechar='"')
        next(reader, None)  # saltar la fila de encabezado
        
        for row in reader:
            dato = {
                'name': row[1],
                'descripcion': row[2],
                'tipo': row[0]
            }
            competencia.insert_one(dato)

            c = competencia.find_one({"name":row[1]})
            c_id = c["_id"]
            pregunta = row[1]+": "+row[2]
            formulario = {
                'idCompetencia': ObjectId(c_id),
                'questions':[
                    {
                        'tipo':'unique',
                        'pregunta': pregunta,
                        'respuestas':[
                            {
                                'puntaje':1,
                                'descripcion':"1"
                            },
                            {
                                'puntaje':2,
                                'descripcion':"2"
                            },
                            {
                                'puntaje':3,
                                'descripcion':"3"
                            },
                            {
                                'puntaje':4,
                                'descripcion':"4"
                            }
                        ]
                    },
                    {
                       'tipo':"texto",
                       'pregunta': "JUSTIFICACIÓN DEL PUNTAJE ASIGNADO",
                        "respuestas": [
                            {
                                "puntaje": 1,
                                "descripcion": "1"
                            }
                        ]
                    }
                ]
            }
            formularioCompetencia.insert_one(formulario)

def populate_cargo(nombre):
    with open(nombre, newline='') as csvfile:
        reader = csv.reader(csvfile, delimiter=',', quotechar='"')
        next(reader, None)  # saltar la fila de encabezado
        
        for row in reader:
            competencias = row[1].split(';')
            lista = []
            for c in competencias:
                comp = competencia.find_one({"name":str(c)})
                c_id = comp['_id']
                lista.append(c_id)

            dato = {
                #'_id': ObjectId(),
                'name': row[0],
                'competencias': [ObjectId(id) for id in lista]
                #'created_at': datetime.datetime.utcnow(),
                #'updated_at': datetime.datetime.utcnow()
            }
            
            cargo.insert_one(dato)

def populate_equipo(nombre):
    equipos_d = {}
    with open(nombre, newline='') as csvfile:
        reader = csv.reader(csvfile, delimiter=',', quotechar='"')
        next(reader, None)  # saltar la fila de encabezado
        
        for row in reader:
            equipos_d[row[0]] = [[],[]]
            dato = {
                'name': row[0],
                'idEvaluador':0,
                'cargos':[]
            }
            
            equipo.insert_one(dato)
    return equipos_d

def populate_usuario(nombre,nombre2,equipos_d):
    with open(nombre, newline='') as csvfile:
        reader = csv.reader(csvfile, delimiter=',', quotechar='"')
        next(reader, None)  # saltar la fila de encabezado
        
        for row in reader:
            u_cargo = cargo.find_one({"name":str(row[4])})
            u_cargo_id = u_cargo['_id']
            u_equipo = equipo.find_one({"name":str(row[5])})
            u_equipo_id = u_equipo['_id']
            equipos_d[row[5]][0].append(u_cargo_id)

            dato = {
                'email':row[0],
                'name': row[2],
                'rol': row[3],
                '_hash': "$2a$04$T55kijSGKWLGVTSc47Wvc.wNmfiVGcHkCyGaLlBzoNVs7UVSAlB7i",
                'cargo':u_cargo_id,
                'team':u_equipo_id
            }
            
            usuario.insert_one(dato)
    with open(nombre2, newline='') as csvfile:
        reader = csv.reader(csvfile, delimiter=',', quotechar='"')
        next(reader, None)  # saltar la fila de encabezado
        
        for row in reader:
            e_user = usuario.find_one({"name":str(row[1])})
            e_user_id = e_user['_id']
            equipos_d[row[0]][1].append(e_user_id)

    for e in equipos_d:
        filter = { 'name': e}
        newvalues = {"$set":{"idEvaluador":equipos_d[e][1][0],"cargos":equipos_d[e][0]}}
        equipo.update_one(filter,newvalues)

populate_competencia('BD_Fusupo-Competencia.csv')
populate_cargo('BD_Fusupo-Cargo-Competencia.csv')
equipo_d1 = populate_equipo('BD_Fusupo-Equipo.csv')
populate_usuario('BD_Fusupo-Usuario.csv','BD_Fusupo-Equipo-Evaluador.csv',equipo_d1)
