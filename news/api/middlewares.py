import jwt
from cryptography.hazmat.backends import default_backend
from cryptography.hazmat.primitives import serialization
from django.http import JsonResponse
from django.conf import settings
from pathlib import Path


class JWTAuthenticationMiddleware:
    def __init__(self, get_response):
        self.get_response = get_response

    def __call__(self, request):
        authorization_header = request.META.get('HTTP_AUTHORIZATION')

        if authorization_header is not None and authorization_header.startswith('Bearer '):
            token = authorization_header.split(' ')[1]
            
            PUBLIC_KEY: Path = settings.BASE_DIR / 'public' / 'public_key.pem'

            with open(PUBLIC_KEY.as_posix(), 'rb') as f:
                public_key = serialization.load_pem_public_key(f.read(), backend=default_backend())

            try:
                payload: dict = jwt.decode(token, public_key, algorithms=['RS256'])
                
                if not isinstance(payload, dict):
                    raise ValueError("Payload not dict instance")

                request.user_data = {
                    "id": payload.get("id", None),
                    "role": payload.get("role", "")
                }
            except jwt.ExpiredSignatureError:
                return JsonResponse({'error': 'Token expired'}, status=401)
            except jwt.InvalidTokenError:
                return JsonResponse({'error': 'Invalid token'}, status=401)
            except ValueError as error:
                return JsonResponse({'error': error}, status=401)

        response = self.get_response(request)
        return response
