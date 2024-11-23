# import pandas as pd
# from typing import List

# def parse_vcf_to_variants(vcf_content: str) -> List[dict]:
#     """Transform VCF data into our variant format."""
#     variants = []
#     for line in vcf_content.split('\n'):
#         if line.startswith('#'):
#             continue
#         if not line:
#             continue
            
#         # Parse VCF fields
#         fields = line.split('\t')
#         if len(fields) < 8:  # VCF has at least 8 columns
#             continue
            
#         variant = {
#             'chromosome': fields[0],
#             'position': int(fields[1]),
#             'reference_allele': fields[3],
#             'alternate_allele': fields[4],
#             'gene_symbol': parse_gene_info(fields[7]),  # Need to parse INFO field
#             'allele_frequency': parse_af_info(fields[7])  # Need to parse INFO field
#         }
#         variants.append(variant)
    
#     return variants