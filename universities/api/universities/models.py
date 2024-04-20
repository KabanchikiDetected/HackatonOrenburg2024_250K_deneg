from django.db import models


class UniversityModel(models.Model):
    name = models.CharField(max_length=256)
    short_name = models.CharField(max_length=64)
    location = models.CharField(max_length=512)
    description = models.TextField(
        blank=True, null=True
    )
    website_url = models.CharField(max_length=256)

    deputy_id = models.CharField(max_length=128, unique=True)
    
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
    user_id = models.CharField(max_length=128)
    group_id = models.IntegerField()
    is_confirmed = models.BooleanField(default=False)
    
    class Meta:
        db_table = "user_to_group"
