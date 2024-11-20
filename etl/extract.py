from pathlib import Path
from utils import load_config

class GnomADExtractor:
    def __init__(self):
        self.config = load_config()
        
    def download_vcf(self, chromosome: str) -> Path:
        """Download VCF file for a specific chromosome"""
        # TODO: Implement download logic
        pass
    
    def extract_vcf_data(self, vcf_path: Path):
        """Extract data from VCF file"""
        # TODO: Implement extraction logic
        pass