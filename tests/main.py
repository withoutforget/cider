import json
import logging
import requests

API_URL = 'http://localhost:8081'

def session_test():
    data = requests.post(
        API_URL + "/api/v1/auth/",
        json = { "username": "admin" },
    )

    if data.status_code != 200:
        logging.error("Status code error %s", data.text)
        return
    
    data:dict = json.loads(data.text)
    logging.info("Got response %s", str(data))
    
    if "error" in data.keys():
        logging.critical("Got error %s", data["error"])
        return

    token: str = data["token"]

    data = requests.post(
        API_URL + "/api/v1/auth/validate/",
        json = { "token": token },
    )

    if data.status_code != 200:
        logging.critical("Status code error %s", data.text)
        return
    
    data:dict = json.loads(data.text)
    logging.info("Got response %s", str(data))
    
    if "error" in data.keys():
        logging.critical("Got error %s", str(data["error"]))
        return

def main():
    logging.basicConfig(level=logging.DEBUG)

    session_test()

if __name__ == '__main__':
    main()