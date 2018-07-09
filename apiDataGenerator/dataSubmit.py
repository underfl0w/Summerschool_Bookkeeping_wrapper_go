import mimesis
from random import *
import requests

n = mimesis

business = n.Business('en')
rand_food = n.Food('fr')
rand_imei = n.Code('nl')
mim_numb = n.Numbers('nl')
mim_pers = n.Person('en')
mim_txt = n.Text('en')
mim_date = n.Datetime('en')



alice_sub_systems = ['Inner Tracking System', 'Time Projection Chamber', ' Time of Flight', ' Photon spectrometer',
                     'Zero Degree Calorimeter', 'ALICE Cosmic Rays Detector', 'Time Projection Chamber',
                     'High Momentum Particle Identification Detector', 'Electro-Magnetic Calorimeter', 'V0 detector',
                     'Transition radiation detector', 'Photon Multiplicity Detector', 'T0 detector',
                     'Forward Multiplicity Detector', 'Muon spectrometer']

alice_class = ['HUMAN', 'MACHINE']

alice_type = ['GENERAL', 'RUN RECORD']
def submit_data(json_data):

    print(json_data)
    url = "http://heikovm.hihva.nl/api/post/entry/data/"
    r = requests.post(url, json=json_data)
    print(r.text)

def generate_data():
    print("data")
    created = mim_date.timestamp(posix=False)
    created = created.replace("T", " ")
    created = created.replace("Z", "")
    sub_system = alice_sub_systems[randint(0, 14)]
    class_alice = alice_class[randint(0, 1)]
    run = str(randint(8000, 80000))
    run_type = alice_type[randint(0, 1)]
    author = mim_pers.full_name()
    title = mim_txt.title()
    log_entry = mim_txt.quote()
    follow_up = mim_txt.quote()
    interruption_duration = mim_date.timestamp(posix=False)
    interruption_duration = interruption_duration.replace("T", " ")
    interruption_duration = interruption_duration.replace("Z", "")
    intervention_type = mim_txt.level()

    test = {"created": ""+created+"", "subsystem": ""+sub_system+"", "class": ""+class_alice+"", "type": ""+run_type+"", "run": ""+run+"", "author": ""+author+"", "title": ""+title+"", "log_entry_text": ""+log_entry+"", "follow_ups": ""+follow_up+"", "interruption_duration": ""+interruption_duration+"", "intervention_type": ""+intervention_type+""}
    submit_data(test)


for x in range (0, 20):
    generate_data()
