from django.db import models


class NewsModel(models.Model):
    title = models.CharField(
        "Название", max_length=512
    )    
    content = models.TextField(
        "Содержимое"
    )
    author_id = models.IntegerField(
        "Id автора"
    )
    