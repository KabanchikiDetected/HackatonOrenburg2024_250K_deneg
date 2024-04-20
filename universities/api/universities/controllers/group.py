from django.core.exceptions import BadRequest
from rest_framework.request import Request
from rest_framework.permissions import AllowAny
from rest_framework.response import Response
from rest_framework.views import APIView
from rest_framework import status, generics
from drf_spectacular.utils import extend_schema, extend_schema_view, OpenApiParameter, OpenApiResponse, OpenApiExample

from .. import docs
from api.permissions import TokenPermission, IsDeputyPermission
from ..services import GroupService
from .. import serializers


class GroupsController(APIView):
    def get_permissions(self):
        if self.request.method in ["GET"]:
            return [AllowAny()]
        
        if self.request.method in ["POST"]:
            return [TokenPermission(), IsDeputyPermission()]
    
    def get(self, request: Request, university_id: int, department_id: int):
        serializer = GroupService.get_all(university_id, department_id)
        
        return Response(
            serializer.data,
            status.HTTP_200_OK
        )

    def post(self, request: Request, university_id: int, department_id: int):
        deputy_id = request.user_data.get("id")
        
        serializer = GroupService.create(university_id, deputy_id, department_id, {
            **request.data,
            "deputy_id": deputy_id,
            "university_id": university_id,
            "department_id": department_id
        })
        
        return Response(
            serializer.data,
            status.HTTP_201_CREATED
        )


class GroupController(APIView):
    def get_permissions(self):
        if self.request.method in ["GET"]:
            return [AllowAny()]
        
        if self.request.method in ["PUT", "DELETE"]:
            return [TokenPermission(), IsDeputyPermission()]

    def get(self, request: Request, university_id: int, department_id: int, group_id: int):
        serializer = GroupService.get_protected(university_id, department_id, group_id)
        
        return Response(
            serializer.data,
            status.HTTP_200_OK
        )
        
    def put(self, request: Request, university_id: int, department_id: int, group_id: int):
        deputy_id = request.user_data.get("id")

        serializer = GroupService.update(university_id, deputy_id, department_id, group_id, {
            **request.data,
            "university_id": university_id,
            "department_id": department_id,
            "deputy_id": deputy_id,
            "group_id": group_id
        })
        
        return Response(
            serializer.data,
            status.HTTP_200_OK
        )
        
    def delete(self, request: Request, university_id: int, department_id: int, group_id: int):
        deputy_id = request.user_data.get("id")
        
        serializer = GroupService.delete(university_id, deputy_id, department_id, group_id)
        
        return Response(
            serializer.data,
            status.HTTP_200_OK
        )


class SearchGroupController(APIView):
    permission_classes = [AllowAny]
    
    def get(self, request: Request):
        search = request.query_params.get("search", "")
        
        serializer = GroupService.search(search)
        
        return Response(
            serializer.data,
            status.HTTP_200_OK
        )


class GroupRequestsController(APIView):
    def get_permissions(self):
        if self.request.method in ["GET"]:
            return [TokenPermission()]
        
        if self.request.method in ["POST"]:
            return [TokenPermission()]
    
    def get(self, request: Request, university_id: int, department_id: int, group_id: int):
        serializer = GroupService.get_requests(university_id, department_id, group_id)
        
        return Response(
            serializer.data,
            status.HTTP_200_OK
        )
        
    def post(self, request: Request, university_id: int, department_id: int, group_id: int):
        user_id = request.user_data.get("id")
        
        serializer = GroupService.create_requests(university_id, department_id, group_id, user_id, {
            **request.data,
            "university_id": university_id,
            "department_id": department_id,
            "group_id": group_id,
            "user_id":group_id
        })
        
        return Response(
            serializer.data,
            status.HTTP_200_OK
        )
        

class GroupRequestAcceptController(APIView):
    permission_classes = [TokenPermission, IsDeputyPermission]
    
    def get(self, request: Request, university_id: int, department_id: int, group_id: int, request_id: int):
        serializer = GroupService.accept_requests(university_id, department_id, group_id, request_id)
        
        return Response(
            serializer.data,
            status.HTTP_200_OK
        )


class GroupRequestDenyController(APIView):
    permission_classes = [TokenPermission, IsDeputyPermission]
    
    def get(self, request: Request, university_id: int, department_id: int, group_id: int, request_id: int):
        serializer = GroupService.deny_requests(university_id, department_id, group_id, request_id)
        
        return Response(
            serializer.data,
            status.HTTP_200_OK
        )


class MyRequestController(APIView):
    permission_classes = [TokenPermission]
    
    def get(self, request: Request):
        user_id = request.user_data.get("id")
        
        serializer = GroupService.get_request_by_user_id(user_id)
        
        return Response(
            serializer.data,
            status.HTTP_200_OK
        )


class MyGroupController(APIView):
    permission_classes = [TokenPermission]
    
    def get(self, request: Request):
        user_id = request.user_data.get("id")

        serializer = GroupService.get_by_user_id(user_id)

        return Response(
            serializer.data,
            status.HTTP_200_OK
        )