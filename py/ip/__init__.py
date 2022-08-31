import requests

def main():
    r = requests.get("https://ifconfig.net/", headers={"accept": "text/plain"})
    print(r.text)
