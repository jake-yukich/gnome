# from sqlalchemy import create_engine
# from sqlalchemy.orm import sessionmaker
# import pandas as pd

# def load_variants_to_db(variants: List[dict], db_url: str):
#     """Load transformed variants into Postgres."""
#     engine = create_engine(db_url)
    
#     # Convert to DataFrame for efficient bulk loading
#     df = pd.DataFrame(variants)
    
#     # Load to database
#     df.to_sql('variants', engine, if_exists='append', index=False)