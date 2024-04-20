import random
from django.db import models


def get_random_number():
    return random.randint(10000, 99999)


class UniversityModel(models.Model):
    name = models.CharField(max_length=256)
    short_name = models.CharField(max_length=64)
    location = models.CharField(max_length=512)
    description = models.TextField(
        blank=True, null=True
    )
    website_url = models.CharField(max_length=256)

    deputy_id = models.CharField(max_length=128, unique=True)
    
    @property
    def images(self):
        images = UniversityImageModel.objects.filter(
            university_id=self.pk
        )

        return images
    
    class Meta:
        db_table = "university"
    
    def __str__(self):
        return f"{self.name} ({self.short_name})"


class DepartmentModel(models.Model):
    name = models.CharField(max_length=256)
    short_name = models.CharField(max_length=64)
    description = models.TextField(
        blank=True, null=True
    )
    
    university = models.ForeignKey(
        UniversityModel,
        on_delete=models.CASCADE, related_name='departments'
    )
    
    class Meta:
        db_table = "department"
    
    def __str__(self):
        return f"{self.name} ({self.short_name})"


class GroupModel(models.Model):
    name = models.CharField(max_length=256)
    department = models.ForeignKey(
        DepartmentModel,
        on_delete=models.CASCADE, related_name='groups'
    )
    
    students_count = models.IntegerField(
        default=0
    )
    
    class Meta:
        db_table = "group"
    
    def __str__(self):
        return self.name


class UserToGroupModel(models.Model):
    user_id = models.CharField(max_length=128, unique=True)
    group_id = models.IntegerField()
    is_confirmed = models.BooleanField(default=False)
    
    class Meta:
        db_table = "user_to_group"
        
    
def university_image_upload_path(instance, filename):
    return f'post_image/{instance.university_id}/{get_random_number()}_{filename}'
        

class UniversityImageModel(models.Model):
    university_id = models.IntegerField(
        "Id университета"
    )
    image = models.ImageField(
        "Изображение",
        upload_to=university_image_upload_path,
        max_length=512,
    )
    
    class Meta:
        db_table = "university_image"
