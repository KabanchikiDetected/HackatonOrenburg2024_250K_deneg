# Generated by Django 4.2.11 on 2024-04-20 09:32

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('universities', '0002_alter_departmentmodel_description_and_more'),
    ]

    operations = [
        migrations.AlterField(
            model_name='groupmodel',
            name='students_count',
            field=models.IntegerField(default=0),
        ),
    ]
