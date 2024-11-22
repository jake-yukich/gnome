from fastapi import FastAPI, Depends
from sqlalchemy.orm import Session
from database import get_db, engine
from variants import Variant, VariantDB

app = FastAPI()

# Create tables on startup
@app.on_event("startup")
async def startup():
    VariantDB.metadata.create_all(bind=engine)

@app.get("/variants/{variant_id}")
def get_variant(variant_id: int, db: Session = Depends(get_db)):
    return db.query(VariantDB).filter(VariantDB.id == variant_id).first()

@app.post("/variants/")
def create_variant(variant: Variant, db: Session = Depends(get_db)):
    db_variant = VariantDB(**variant.dict())
    db.add(db_variant)
    db.commit()
    db.refresh(db_variant)
    return db_variant