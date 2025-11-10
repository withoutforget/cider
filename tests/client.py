import requests
from typing import Optional
from api_types import RegisterUserResponse, CreateSessionResponse, RevokeSessionResponse, ValidateSessionResponse
from adaptix import Retort

class CiderAPI:
    def __init__(self, host: str = 'localhost', port: int = 8081):
        self.base_url = f'http://{host}:{port}'
        self.prefix = '/api/v1'
        self.session = requests.Session()
        self.token: Optional[str] = None
        self.retort = Retort()
    
    def _url(self, endpoint: str) -> str:
        return f'{self.base_url}{self.prefix}{endpoint}'
    
    def _post(self, endpoint: str, data: dict, expected: int = 200) -> dict:
        resp = self.session.post(self._url(endpoint), json=data, timeout=5)
        assert resp.status_code == expected, f"Expected {expected}, got {resp.status_code}: {resp.text}"
        return resp.json() if resp.text else {}
    
    def register(self, username: str, password: str) -> RegisterUserResponse:
        data = self._post('/auth/register', 
                         {'username': username, 'password': password}, 
                         expected=201)
        return self.retort.load(data, RegisterUserResponse)
    
    def login(self, username: str, password: str) -> CreateSessionResponse:
        data = self._post('/auth/', {'username': username, 'password': password})
        resp = self.retort.load(data, CreateSessionResponse)
        self.token = resp.token
        return resp
    
    def validate(self, token: str = None) -> ValidateSessionResponse:
        token = token or self.token
        assert token, "No token to validate"
        data = self._post('/auth/validate/', {'token': token})
        resp =  self.retort.load(data, ValidateSessionResponse)
        return resp
    
    def revoke(self, token: str = None) -> RevokeSessionResponse:
        token = token or self.token
        assert token, "No token to validate"
        data = self._post('/auth/revoke/', {'token': token})
        resp =  self.retort.load(data, RevokeSessionResponse)
        return resp