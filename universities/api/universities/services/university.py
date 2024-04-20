from django.core.exceptions import BadRequest
from django.db.models import Q
from rest_framework.serializers import ListSerializer

from .. import models
from .. import serializers


class UniversityService:
    @staticmethod
    def create(deputy_id: str, university_data: dict):
        serializer = serializers.UniversitySerializer(data={
            **university_data,
            "deputy_id": deputy_id
        })
        
        if not serializer.is_valid():
            raise BadRequest(serializer.errors)
        
        serializer.save()
        
        return serializer
    
    @staticmethod
    def update(university_id: int, deputy_id: str, new_university_data: dict) -> serializers.UniversitySerializer:
        university = UniversityService._get_by_id_and_deputy_id(university_id, deputy_id)
        
        serializer = serializers.UniversitySerializer(university, data=new_university_data)
        
        if not serializer.is_valid():
            raise BadRequest(serializer.errors)
        
        serializer.save()
    
        return serializer
    
    @staticmethod
    def delete(university_id: int, deputy_id: str):
        university = UniversityService._get_by_id_and_deputy_id(university_id, deputy_id)
        
        university.delete()
        
        serializer = serializers.EmptySerializer()
        serializer.data = "University deleted"

        return serializer

    @staticmethod
    def get_all():
        universities = models.UniversityModel.objects.all()
        
        serializer = serializers.UniversitySerializer(universities, many=True)
        
        return serializer
    
    @staticmethod
    def add_image(university_id: int, image_data: dict):
        university = UniversityService._get_by_id(university_id)
        
        serializer = serializers.UniversityImageSerialzier(data=image_data)

        if not serializer.is_valid():
            raise BadRequest(serializer.errors)
        
        serializer.save()
        
        serializer = serializers.EmptySerializer()
        serializer.data = "Image added"
        
        return serializer
    
    @staticmethod
    def get_one_by_id(university_id: int):
        university = UniversityService._get_by_id(university_id)
        
        serializer = serializers.UniversitySerializer(university)
        
        return serializer
    
    @staticmethod
    def get_one_by_deputy_id(deputy_id: str):
        try:
            university = models.UniversityModel.objects.get(
                deputy_id=deputy_id
            )
        except models.UniversityModel.DoesNotExist:
            raise BadRequest(f"Not found university with deputy_id {deputy_id}")
        
        serializer = serializers.UniversitySerializer(university)
        
        return serializer
    
    @staticmethod
    def search_by_name(search: str):
        universities = models.UniversityModel.objects.filter(
            Q(name__icontains=search) | Q(short_name__icontains=search)
        ).distinct()
        
        serializer = serializers.UniversitySerializer(universities, many=True)
        
        return serializer
    
    @staticmethod
    def _get_by_id(university_id: int):
        try:
            university = models.UniversityModel.objects.get(
                pk=university_id,
            )
        except models.UniversityModel.DoesNotExist:
            raise BadRequest(f"Not found university with id {university_id}")
        
        return university
    
    @staticmethod
    def _get_by_id_and_deputy_id(university_id: int, deputy_id: str):
        try:
            university = models.UniversityModel.objects.get(
                pk=university_id,
                deputy_id=deputy_id
            )
        except models.UniversityModel.DoesNotExist:
            raise BadRequest(f"Not found university with id {university_id} by deputy with id {deputy_id}")
        
        return university
