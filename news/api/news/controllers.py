from typing import Any
from rest_framework.request import Request
from rest_framework.response import Response
from rest_framework.views import APIView
from rest_framework import status


class BaseController(APIView):
    def perform_authentication(self, request):
        return None


class NewsController(BaseController):
    def get(self, request: Request):

        return Response({
            "message": "ok"
        }, status.HTTP_200_OK)

    def post(self, request: Request):
        ...
