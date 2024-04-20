from rest_framework.permissions import BasePermission
from rest_framework.request import Request


class TokenPermission(BasePermission):
    def has_permission(self, request: Request, view):
        return isinstance(getattr(request, "user_data", None), dict)
