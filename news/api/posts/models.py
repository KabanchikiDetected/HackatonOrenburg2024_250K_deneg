from django.db import models
import random
import time


class UnixTimestampManager(models.Manager):
    def create(self, **kwargs):
        kwargs['created_at'] = str(time.time())
        return super().create(**kwargs)
    

def get_random_number():
    return random.randint(10000, 99999)


class PostModel(models.Model):
    title = models.CharField(
        "Название", max_length=512
    )    
    content = models.TextField(
        "Содержимое"
    )
    author_id = models.CharField(
        "Id автора"
    )

    likes = models.IntegerField(
        "Кол-во лайков", default=0
    )
    
    created_at = models.CharField(
        "Дата создания",
        blank=True, null=True
    )
    
    objects = UnixTimestampManager()

    class Meta:
        db_table = "post"
        verbose_name = "Пост"

    @property
    def images(self):
        images = PostImageModel.objects.filter(
            post_id=self.pk
        )
        
        return images


class LikesToPostModel(models.Model):
    user_id = models.CharField(
        "Id пользователя"
    )
    post_id = models.IntegerField(
        "Id поста"
    )
    
    class Meta:
        db_table = "like_to_post"
        verbose_name = "Лайки к постам"


def post_image_upload_path(instance, filename):
    return f'post_image/{instance.post_id}/{get_random_number()}_{filename}'


class PostImageModel(models.Model):
    post_id = models.IntegerField(
        "Id поста"
    )
    image = models.ImageField(
        "Изображение",
        upload_to=post_image_upload_path,
        max_length=512,
    )
    
    class Meta:
        db_table = "post_image"
        verbose_name = "Лайки к постам"
