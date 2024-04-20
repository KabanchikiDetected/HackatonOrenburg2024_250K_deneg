import aiohttp
import asyncio

from config import SEARCH_URLS


async def filter_service(search):
    # async with aiohttp.ClientSession() as session:
    #     result = {}
    #     for name, url in SEARCH_URLS.items():
    #         async with session.get(url, params={ "search": search }) as response:
    #             result[name] = await response.json()
                
    #     return result
    return []
