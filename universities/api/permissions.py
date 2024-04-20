from rest_framework.permissions import BasePermission
from rest_framework.request import Request


class TokenPermission(BasePermission):
    def has_permission(self, request: Request, view):
        return isinstance(getattr(request, "user_data", None), dict)


class IsDeputyPermission(BasePermission):
    def has_permission(self, request: Request, view):
        user_data = getattr(request, "user_data", {})
        
        print(user_data)
        
        return user_data.get("role") == "deputy"
