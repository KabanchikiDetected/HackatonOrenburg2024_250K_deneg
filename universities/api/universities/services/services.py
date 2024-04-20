from .department import DepartmentService
from .group import GroupService

from .. import serializers


def get_full_university_by_user_id(user_id: str):
    group = GroupService._get_by_user_id(user_id)
    department = group.department
    university = department.university
    
    
    serializer = serializers.EmptySerializer()
    serializer.data = {
        "university": serializers.UniversitySerializer(university).data,
        "department": serializers.DepartmentSerializer(department).data,
        "group": serializers.GroupSerializer(group).data
    }

    return serializer
