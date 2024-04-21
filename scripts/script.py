import random
import requests
import json

tokens = []
ids_for_request_add_student = []


def create_users():
    users = json.load(open("users.json"))

    for user in users:
        response = requests.post("http://127.0.0.1/api/auth/register", json=user)
        if response.status_code != 201:
            print(f"Failed to create user {user['email']} with status code {response.status_code}")
            continue

        # Login and get token
        data = {
            "email": user["email"],
            "password": user["password"]
        }
        response = requests.post("http://127.0.0.1/api/auth/login", json=data)
        token = response.json().get("token")
        print(f"Created user {data['email']}")
        tokens.append(token)

def make_deputy(token):
    response = requests.post("http://127.0.0.1/api/auth/make_deputy", headers={
        "Authorization": f"Bearer {token}"
    })
    return response.json()["token"]


def make_universities():
    universities = json.load(open("universities.json"))
    facults = json.load(open("facults.json"))
    cafedras = json.load(open("cafedras.json"))
    for i, university in enumerate(universities):
        token = make_deputy(tokens[i])
        response = requests.post("http://127.0.0.1/api/universities/university/",
                        headers={
                            "Authorization": f"Bearer {token}"
                            },
                        json=university
                        )
        if response.status_code != 200 and response.status_code != 201:
            print(f"Failed to create university {university['name']} with status code {response.status_code}")
            print(university)
            exit()
        university_id = int(response.json()["id"])
        print(f"Created university {university['name']} with id {university_id}")

        for j, facult in enumerate(facults[i]):
            response = requests.post(f"http://127.0.0.1/api/universities/university/{university_id}/department/",
                            headers={
                                "Authorization": f"Bearer {token}"
                                },
                            json=facult
                            )
            if response.status_code != 200 and response.status_code != 201:
                print(f"\tFailed to create department {facult['name']} with status code {response.status_code}")
                print(universities[i])
                exit()
            department_id = int(response.json()["id"])
            print(f"\tCreated department {facult['name']} with id {department_id}")

            for cafedra in cafedras[i][j]:
                response = requests.post(f"http://127.0.0.1/api/universities/university/{university_id}/department/{department_id}/group/",
                                headers={
                                    "Authorization": f"Bearer {token}"
                                    },
                                json=cafedra
                                )
                if response.status_code != 200 and response.status_code != 201:
                    print(f"\t\tFailed to create cafedra {cafedra['name']} with status code {response.status_code}")
                    print(universities[i])
                    exit()
                cafedra_id = int(response.json()["id"])
                print(f"\t\tCreated cafedra {cafedra['name']} with id {cafedra_id}")
                ids_for_request_add_student.append((token, university_id, department_id, cafedra_id))

def make_students():
    students: list[dict] = json.load(open("students.json"))
    for i, student in enumerate(students):
        deputy_token,u_id, d_id, g_id = random.choice(ids_for_request_add_student)
        
        # Create student
        student["faculty_id"] = str(g_id)
        response = requests.post(f"http://127.0.0.1/api/students/",
                        headers={
                            "Authorization": f"Bearer {tokens[i+2]}" # two first users is deputy
                        },
                        json=student
                        )
        if response.status_code != 200 and response.status_code != 201:
            print("==========================================================")
            print(f"Failed to create student {student['first_name'] + ' ' + student['last_name']} with status code {response.status_code}")
            # print(response.status_code)
            print(response.content)
            print(student)
            exit()
        print(f"Created student {response.json()['name']}")

        # Add photo
        response = requests.post(f"http://127.0.0.1/api/students/me/image)",
                      headers={
                          "Authorization": f"Bearer {tokens[i+2]}",
                          "Content-Type": "multipart/form-data",
                      },
                      files={"image": (open(f"./images/{student['first_name']}{student['last_name']}.jpg", "rb"))})
        
        if response.status_code != 200 and response.status_code != 201:
            # print(tokens[i+2])
            print(f"Failed to add photo with status code {response.status_code}")
            print(response.content)
            # exit()

        # Create request
        response = requests.post(f"http://127.0.0.1/api/universities/university/{u_id}/department/{d_id}/group/{g_id}/requests/",
                                 headers={
                                     "Authorization": f"Bearer {tokens[i+2]}"
                                 })
        if response.status_code != 200 and response.status_code != 201:
            print(f"Failed to create request with status code {response.status_code}")
            print(response.json())
            exit()
        print(f"Created request with id {response.json()['id']}")

        # accept request
        request_id = int(response.json()["id"])
        deputy_token, u_id, d_id, g_id = random.choice(ids_for_request_add_student)
        response = requests.get(f"http://127.0.0.1/api/universities/university/{u_id}/department/{d_id}/group/{g_id}/requests/{request_id}/accept/",
                                 headers={
                                     "Authorization": f"Bearer {deputy_token}"
                                 })
        if response.status_code != 200 and response.status_code != 201:
            print(f"Failed to accept request with status code {response.status_code}")
            print(response.json())
            exit()
        print(f"Accepted request with id {request_id}")

        
def main():
    create_users()
    print("\n")
    make_universities()
    print("\n")
    make_students()
    


if __name__ == '__main__':
    main()