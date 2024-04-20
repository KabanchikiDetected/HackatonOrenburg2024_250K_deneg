from fastapi import APIRouter, Request, Depends

from .service import search_service


router = APIRouter(
    prefix="/search"
)

@router.get("")
async def search_controller(search: str):
    results = await search_service(search)
    
    return results
