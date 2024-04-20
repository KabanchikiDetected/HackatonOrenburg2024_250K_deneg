from django.core.exceptions import BadRequest
from rest_framework.serializers import ListSerializer

from .. import models
from .. import serializers
from .department import DepartmentService, UniversityService


class GroupService:
    @staticmethod
    def create(university_id: int, deputy_id: str, department_id: int, group_data: dict):
        university = UniversityService._get_by_id_and_deputy_id(university_id, deputy_id)
        department = DepartmentService._get_by_university_id_and_id(university.pk, department_id)
        
        serializer = serializers.GroupSerializer(data={
            **group_data,
            "department": department.pk
        })
        
        if not serializer.is_valid():
            raise BadRequest(serializer.errors)
        
        serializer.save(department=department)
        
        return serializer
    
    @staticmethod
    def get_all(university_id: int, department_id: int):
        department = DepartmentService._get_by_university_id_and_id(
            university_id, department_id
        )

        groups = models.GroupModel.objects.filter(
            department_id=department.pk
        )
        
        serializer = serializers.GroupSerializer(groups, many=True)
        
        return serializer
    
    @staticmethod
    def update(university_id: int, deputy_id: str, department_id: int, group_id: int, new_department_data):
        university = UniversityService._get_by_id_and_deputy_id(university_id, deputy_id)
        department = DepartmentService._get_by_university_id_and_id(university.pk, department_id)
        group = GroupService._get_protected(university.pk, department.pk, group_id)
        
        serializer = serializers.GroupSerializer(group, {
            **new_department_data,
            "department": department.pk
        })
        
        if not serializer.is_valid():
            raise BadRequest(serializer.errors)
        
        serializer.save()
        
        return serializer
    
    @staticmethod
    def delete(university_id: int, deputy_id: str, department_id: int, group_id: int):
        university = UniversityService._get_by_id_and_deputy_id(university_id, deputy_id)
        department = DepartmentService._get_by_university_id_and_id(university.pk, department_id)
        group = GroupService._get_protected(university.pk, department.pk, group_id)
        
        group.delete()
        
        serializer = serializers.EmptySerializer()
        serializer.data = "Group deleted"
        
        return serializer
    
    @staticmethod
    def search(search: str):
        groups = models.GroupModel.objects.filter(
            name__icontains=search
        )

        serializer = serializers.GroupSerializer(groups, many=True)
        
        return serializer
    
    @staticmethod
    def get_protected(university_id: int, department_id: int, group_id: int):
        group = GroupService._get_protected(university_id, department_id, group_id)
        
        serializer = serializers.GroupSerializer(group)
        
        return serializer

    @staticmethod
    def _get_protected(university_id: int, department_id: int, group_id: int):
        try:
            department = DepartmentService._get_by_university_id_and_id(university_id, department_id)
            
            group = models.GroupModel.objects.get(
                pk=group_id,
                department_id=department.pk,
            )
        except models.GroupModel.DoesNotExist:
            raise BadRequest(f"Not found group with id {group_id} by department with id {department_id} and university with id {university_id}")

        return group

    @staticmethod
    def _get_by_id(group_id: int):
        try:
            group = models.GroupModel.objects.get(
                pk=group_id
            )
        except models.GroupModel.DoesNotExist:
            raise BadRequest(f"Not found group with id {group}")
        
        return group
    
    @staticmethod
    def create_requests(university_id: int, department_id: int, group_id: int, user_id: int, request_data: dict):
        group = GroupService._get_protected(university_id, department_id, group_id)
        
        
        serializer = serializers.UserToGroupSerializer(data={
            "user_id": user_id,
            "group_id": group.id,
            "is_confirmed": False
        })
        
        if not serializer.is_valid():
            raise BadRequest(serializer.errors)

        serializer.save()

        return serializer
    
    @staticmethod
    def get_requests(university_id: int, department_id: int, group_id: int):
        users_to_group = GroupService._get_requests(university_id, department_id, group_id)
        
        serializer = serializers.UserToGroupSerializer(users_to_group, many=True)

        return serializer

    @staticmethod
    def _get_requests(university_id: int, department_id: int, group_id: int):
        group = GroupService._get_protected(university_id, department_id, group_id)
        
        users_to_group = models.UserToGroupModel.objects.filter(
            group_id=group.pk,
            is_confirmed=False
        )
        
        return users_to_group
    
    @staticmethod
    def accept_requests(university_id: int, department_id: int, group_id: int, request_id: int):
        group = GroupService._get_protected(university_id, department_id, group_id)
        
        users_to_group = GroupService._get_request(request_id)
        users_to_group.is_confirmed = True
        users_to_group.save()
        
        group.students_count += 1
        group.save()
        
        serializer = serializers.EmptySerializer()
        serializer.data = "Request accepted"

        return serializer
    
    @staticmethod
    def deny_requests(university_id: int, department_id: int, group_id: int, request_id: int):
        group = GroupService._get_protected(university_id, department_id, group_id)
        
        users_to_group = GroupService._get_request(request_id)
        users_to_group.delete()
        
        serializer = serializers.EmptySerializer()
        serializer.data = "Request deny"

        return serializer
    
    @staticmethod
    def _get_request(request_id: int):
        try:
            request = models.UserToGroupModel.objects.get(
                pk=request_id
            )
        except models.UserToGroupModel.DoesNotExist:
            raise BadRequest(f"No reqeust with id {request_id}")

        return request
    
    @staticmethod
    def get_request_by_user_id(user_id: str):
        group_to_user = GroupService._get_request_by_user_id(user_id)
        
        serializer = serializers.UserToGroupSerializer(group_to_user)
        
        return serializer
    
    @staticmethod
    def _get_request_by_user_id(user_id: str):
        try:
            request = models.UserToGroupModel.objects.get(
                user_id=user_id
            )
        except models.UserToGroupModel.DoesNotExist:
            raise BadRequest(f"You havent requests to group")

        return request
    
    @staticmethod
    def get_by_user_id(user_id: str):
        group = GroupService._get_by_user_id(user_id)
        
        serializer = serializers.GroupSerializer(group)
        
        return serializer
        
    @staticmethod
    def _get_by_user_id(user_id: str):
        group_to_user = GroupService._get_request_by_user_id(user_id)
        
        if not group_to_user.is_confirmed:
            serializer = serializers.EmptySerializer()
            serializer.data = "You dont have group"
            
            return serializer
        
        
        group = GroupService._get_by_id(group_to_user.group_id)
        
        return group
