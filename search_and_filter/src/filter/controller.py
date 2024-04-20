from fastapi import APIRouter, Request, Depends

from .service import filter_service


router = APIRouter(
    prefix="/filter"
)

@router.get("")
async def filter_controller(search: str):
    results = await filter_service(search)
    
    return results
