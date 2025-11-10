from dataclasses import dataclass
from typing import Optional


@dataclass(slots=True)
class CreateSessionResponse:
    token: Optional[str] = None
    error: Optional[str] = None


@dataclass(slots=True)
class RegisterUserResponse:
    user_id: Optional[int] = None
    error: Optional[str] = None


@dataclass(slots=True)
class SessionModel:
    user_id: int
    username: str
    created_at: Optional[str] = None
    expires_at: Optional[str] = None


@dataclass(slots=True)
class ValidateSessionResponse:
    session: Optional[SessionModel] = None
    error: Optional[str] = None