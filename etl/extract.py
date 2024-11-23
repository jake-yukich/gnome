# import requests
# import gzip
# import io

# def download_gnomad_data(chromosome: str) -> str:
#     """Download gnomAD VCF for a specific chromosome."""
#     # Example URL (we'll need to find exact gnomAD URLs)
#     url = f"https://gnomad.broadinstitute.org/download/vcf/chr{chromosome}.vcf.gz"
    
#     response = requests.get(url, stream=True)
#     if response.status_code == 200:
#         return gzip.decompress(response.content).decode('utf-8')
#     else:
#         raise Exception(f"Failed to download data: {response.status_code}")