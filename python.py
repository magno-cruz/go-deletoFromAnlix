import requests
from requests.auth import HTTPBasicAuth
'''
username = "magno"
password = "10203040"
url = '168.205.100.82/api/v2/device/update/FHTT94087C20'

response = requests.get(url, auth=(username, password))
print(response.status_code)
print(response.json())
'''

response = requests.get('https://api.github.com / user, ', auth = HTTPBasicAuth('user', 'pass'))
print(response)