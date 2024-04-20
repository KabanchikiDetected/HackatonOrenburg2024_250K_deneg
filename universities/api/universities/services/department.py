from django.core.exceptions import BadRequest
from django.db.models import Q
from rest_framework.serializers import ListSerializer

from .. import models
from .. import serializers
from .university import UniversityService


class DepartmentService:
    @staticmethod
    def create(university_id: int, deputy_id: str, department_data: dict):
        university = UniversityService._get_by_id_and_deputy_id(university_id, deputy_id)
        
        serializer = serializers.DepartmentSerializer(data={
            **department_data,
            "university": university.pk
        })
        
        if not serializer.is_valid():
            raise BadRequest(serializer.errors)
        
        serializer.save(university=university)
        
        return serializer
    
    @staticmethod
    def get_all(university_id: int):
        departments = DepartmentService._get_all(university_id)
        
        serializer = serializers.DepartmentSerializer(departments, many=True)
        
        return serializer
    
    @staticmethod
    def update(university_id: int, deputy_id: str, department_id: int, new_department_data):
        university = UniversityService._get_by_id_and_deputy_id(university_id, deputy_id)
        department = DepartmentService._get_by_university_id_and_id(university.pk, department_id)
        
        serializer = serializers.DepartmentSerializer(department, {
            **new_department_data,
            "university": university.pk
        })
        
        if not serializer.is_valid():
            raise BadRequest(serializer.errors)
        
        serializer.save()
        
        return serializer
    
    @staticmethod
    def delete(university_id: int, deputy_id: str, department_id: int):
        university = UniversityService._get_by_id_and_deputy_id(university_id, deputy_id)
        department = DepartmentService._get_by_university_id_and_id(university.pk, department_id)
        
        department.delete()
        
        serializer = serializers.EmptySerializer()
        serializer.data = "Department deleted"
        
        return serializer
    
    @staticmethod
    def search(search: str):
        departments = models.DepartmentModel.objects.filter(
            Q(name__icontains=search) | Q(short_name__icontains=search)
        )

        serializer = serializers.DepartmentSerializer(departments, many=True)
        
        return serializer
    
    @staticmethod
    def _get_all(university_id):
        university = UniversityService._get_by_id(university_id)
        
        departments = models.DepartmentModel.objects.filter(
            university_id=university.pk
        )
        
        return departments
    
    @staticmethod
    def get_by_university_id_and_id(university_id: int, department_id: int):
        department = DepartmentService._get_by_university_id_and_id(university_id, department_id)
        
        serializer = serializers.DepartmentSerializer(department)
        
        return serializer

    @staticmethod
    def _get_by_university_id_and_id(university_id: int, department_id: int):
        try:
            department = models.DepartmentModel.objects.get(
                pk=department_id,
                university_id=university_id
            )
        except models.DepartmentModel.DoesNotExist:
            raise BadRequest(f"Not found department with id {department_id} by university with id {university_id}")
        
        return department
    
    @staticmethod
    def _get_by_id(department_id: int):
        try:
            department = models.DepartmentModel.objects.get(
                pk=department_id
            )
        except models.DepartmentModel.DoesNotExist:
            raise BadRequest(f"Not found department with id {department_id}")
        
        return department
