from sqlalchemy import Column, Integer, String, Float
from pydantic import BaseModel
from database import Base

# Database Model
class VariantDB(Base):
    __tablename__ = "variants"
    
    id = Column(Integer, primary_key=True)
    chromosome = Column(String)
    position = Column(Integer)
    reference_allele = Column(String)
    alternate_allele = Column(String)
    gene_symbol = Column(String)
    allele_frequency = Column(Float)

# API Model
class Variant(BaseModel):
    chromosome: str
    position: int
    reference_allele: str
    alternate_allele: str
    gene_symbol: str
    allele_frequency: float

    class Config:
        from_attributes = True  # Allows conversion from SQLAlchemy model