import requests

def get_ip():
    r = requests.get("https://ifconfig.net/", headers={"accept": "text/plain"})
    return r.text
