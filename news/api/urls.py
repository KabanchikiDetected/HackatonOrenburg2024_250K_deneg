from django.contrib import admin
from django.urls import path, include
from drf_spectacular.views import SpectacularAPIView, SpectacularSwaggerView


urlpatterns = [
    path('news/', include("api.posts.urls")),
    
    path('news/schema/', SpectacularAPIView.as_view(), name='schema'),
    path('news/docs/', SpectacularSwaggerView.as_view(url_name='schema'), name='docs')
]
