from rest_framework.request import Request
from rest_framework.permissions import AllowAny
from rest_framework.response import Response
from rest_framework.views import APIView
from rest_framework import status, generics
from drf_spectacular.utils import extend_schema, extend_schema_view, OpenApiParameter, OpenApiResponse, OpenApiExample

from .. import docs
from api.permissions import IsDeputyPermission, TokenPermission
from ..services import UniversityService, get_full_university_by_user_id
from .. import serializers


class UniversitiesController(APIView):
    def get_permissions(self):
        if self.request.method in ["GET"]:
            return [AllowAny()]
        
        if self.request.method in ["POST"]:
            return [TokenPermission(), IsDeputyPermission()]
    
    def get(self, request: Request):
        serializer = UniversityService.get_all()
        
        return Response(
            serializer.data,
            status.HTTP_200_OK
        )

    def post(self, request: Request):
        deputy_id = request.user_data.get("id")
        
        serializer = UniversityService.create(deputy_id, {
            **request.data,
            "deputy_id": deputy_id
        })
        
        return Response(
            serializer.data,
            status.HTTP_201_CREATED
        )


class UniversityController(APIView):
    def get_permissions(self):
        if self.request.method in ["GET"]:
            return [AllowAny()]
        
        if self.request.method in ["PUT", "DELETE"]:
            return [TokenPermission(), IsDeputyPermission()]
    
    def get(self, request: Request, university_id: int):
        serializer = UniversityService.get_one_by_id(university_id)
        
        return Response(
            serializer.data,
            status.HTTP_200_OK
        )

    def put(self, request: Request, university_id: int):
        deputy_id = request.user_data.get("id")
        
        serializer = UniversityService.update(university_id, deputy_id, {
            **request.data,
            "id": university_id,
            "deputy_id": deputy_id
        })

        return Response(
            serializer.data,
            status.HTTP_201_CREATED
        )
        
    def delete(self, request: Request, university_id: int):
        deputy_id = request.user_data.get("id")
        
        serializer = UniversityService.delete(university_id, deputy_id)
        
        return Response(
            serializer.data,
            status.HTTP_200_OK
        )


class UniversitySearchController(APIView):
    permission_classes = [AllowAny]
    
    def get(self, request: Request):
        search = request.query_params.get("search", "")
        
        serializer = UniversityService.search_by_name(search)
        
        return Response(
            serializer.data,
            status.HTTP_200_OK
        )


class DeputyUniversityController(APIView):
    permission_classes = [TokenPermission, IsDeputyPermission]
    
    def get(self, request: Request):
        deputy_id = request.user_data.get("id")
        
        serializer = UniversityService.get_one_by_deputy_id(deputy_id)
        
        return Response(
            serializer.data,
            status.HTTP_200_OK
        )


class UniversityAddImageController(APIView):
    permission_classes = [TokenPermission, IsDeputyPermission]
    
    def post(self, request: Request, university_id: int):
        serializer = UniversityService.add_image(university_id, {
            "image": request.FILES.get("image"),
            "university_id": university_id,
        })
        
        return Response(
            serializer.data,
            status.HTTP_201_CREATED
        )


class MyEducationController(APIView):
    def get(self, request: Request):
        user_id = request.user_data.get("id")
        
        serializer = get_full_university_by_user_id(user_id)
        
        return Response(
            serializer.data,
            status.HTTP_200_OK
        )


class UserEducationController(APIView):
    permission_classes = [AllowAny]
    
    def get(self, request: Request, user_id: str):        
        serializer = get_full_university_by_user_id(user_id)
        
        return Response(
            serializer.data,
            status.HTTP_200_OK
        )
