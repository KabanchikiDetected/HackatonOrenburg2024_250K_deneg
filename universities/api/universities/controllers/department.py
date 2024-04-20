from django.core.exceptions import BadRequest
from rest_framework.request import Request
from rest_framework.permissions import AllowAny
from rest_framework.response import Response
from rest_framework.views import APIView
from rest_framework import status, generics
from drf_spectacular.utils import extend_schema, extend_schema_view, OpenApiParameter, OpenApiResponse, OpenApiExample

from .. import docs
from api.permissions import TokenPermission, IsDeputyPermission
from ..services import DepartmentService
from .. import serializers


class DepartmentsController(APIView):
    def get_permissions(self):
        if self.request.method in ["GET"]:
            return [AllowAny()]
        
        if self.request.method in ["POST"]:
            return [TokenPermission(), IsDeputyPermission()]
    
    def get(self, request: Request, university_id: int):
        serializer = DepartmentService.get_all(university_id)
        
        return Response(
            serializer.data,
            status.HTTP_200_OK
        )

    def post(self, request: Request, university_id: int):
        deputy_id = request.user_data.get("id")
        
        serializer = DepartmentService.create(university_id, deputy_id, {
            **request.data,
            "deputy_id": deputy_id,
            "university_id": university_id
        })
        
        return Response(
            serializer.data,
            status.HTTP_201_CREATED
        )


class DepartmentController(APIView):
    def get_permissions(self):
        if self.request.method in ["GET"]:
            return [AllowAny()]
        
        if self.request.method in ["PUT", "DELETE"]:
            return [TokenPermission(), IsDeputyPermission()]

    def get(self, request: Request, university_id: int, department_id: int):
        serializer = DepartmentService.get_by_university_id_and_id(university_id, department_id)
        
        return Response(
            serializer.data,
            status.HTTP_200_OK
        )
        
    def put(self, request: Request, university_id: int, department_id: int):
        deputy_id = request.user_data.get("id")

        serializer = DepartmentService.update(university_id, deputy_id, department_id, {
            **request.data,
            "university_id": university_id,
            "department_id": department_id,
            "deputy_id": deputy_id
        })
        
        return Response(
            serializer.data,
            status.HTTP_200_OK
        )
        
    def delete(self, request: Request, university_id: int, department_id: int):
        deputy_id = request.user_data.get("id")
        
        serializer = DepartmentService.delete(university_id, deputy_id, department_id)
        
        return Response(
            serializer.data,
            status.HTTP_200_OK
        )


class SearchDepartmentController(APIView):
    permission_classes = [AllowAny]

    def get(self, request: Request):
        search = request.query_params.get("search", "")
        
        serializer = DepartmentService.search(search)
        
        return Response(
            serializer.data,
            status.HTTP_200_OK
        )
