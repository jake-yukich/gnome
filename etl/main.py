# from src.extract import download_gnomad_data
# from src.transform import parse_vcf_to_variants
# from src.load import load_variants_to_db
# import os

# def run_pipeline(chromosome: str):
#     # Extract
#     vcf_content = download_gnomad_data(chromosome)
    
#     # Transform
#     variants = parse_vcf_to_variants(vcf_content)
    
#     # Load
#     db_url = os.getenv("DATABASE_URL", "postgresql://postgres:postgres@localhost:5432/gnome")
#     load_variants_to_db(variants, db_url)

# if __name__ == "__main__":
#     # Run for chromosome 1 as test
#     run_pipeline("1")