import string
import requests

url=""
data={"cmd":"ls"}
requests.post(url,data,timeout=1)