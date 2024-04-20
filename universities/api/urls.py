from django.contrib import admin
from django.urls import path, include
from drf_spectacular.views import SpectacularAPIView, SpectacularSwaggerView


urlpatterns = [
    path('university/', include("api.universities.urls")),
    
    path('university/schema/', SpectacularAPIView.as_view(), name='schema'),
    path('university/docs/', SpectacularSwaggerView.as_view(url_name='schema'), name='docs')
]
