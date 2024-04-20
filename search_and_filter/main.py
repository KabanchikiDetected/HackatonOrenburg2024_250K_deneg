import os
import uvicorn
from fastapi import FastAPI, APIRouter
from dotenv import load_dotenv

from src.search import controller as search_controller


load_dotenv()

app = FastAPI()

api_router = APIRouter(
    prefix="/api/search-and-filter"
)
api_router.include_router(search_controller.router)

app.include_router(api_router)


@app.get("/")
async def index():
    return "Filter and Search Microservice"


def start():
    uvicorn.run(
        "main:app",
        host=os.getenv("HOST", "127.0.0.1"),
        port=int(os.getenv("PORT", 8001)),
        reload=True,
        workers=1
    )
