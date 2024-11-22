from fastapi import FastAPI
from variants import Variant

app = FastAPI()

# @app.get("/variants/{variant_id}")
# async def get_variant(variant_id: int):
#     # Get variant from database
#     return {"message": "Found variant"}

# @app.post("/variants/analyze")
# async def analyze_variant(variant: Variant):
#     # Analyze variant with LLM
#     return {"analysis": "Variant analysis"}