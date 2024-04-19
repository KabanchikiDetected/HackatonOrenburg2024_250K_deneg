from rest_framework.request import Request
from rest_framework.response import Response
from rest_framework.views import APIView


class NewsController(APIView):
    def get(self, request: Request):
        ...

    def post(self, request: Request):
        ...
