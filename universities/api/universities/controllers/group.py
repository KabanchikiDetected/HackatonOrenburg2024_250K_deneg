from django.core.exceptions import BadRequest
from rest_framework.request import Request
from rest_framework import permissions
from rest_framework.response import Response
from rest_framework.views import APIView
from rest_framework import status, generics
from drf_spectacular.utils import extend_schema, extend_schema_view, OpenApiParameter, OpenApiResponse, OpenApiExample

from .. import docs
from ..services import GroupService
from .. import serializers


class GroupController(APIView):
    ...
