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
