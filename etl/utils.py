import yaml
from pathlib import Path

def load_config(config_path: str = None) -> dict:
    """Load configuration from yaml file"""
    if config_path is None:
        config_path = Path(__file__).parent / 'config.yaml'
    
    with open(config_path, 'r') as f:
        return yaml.safe_load(f)

def setup_logging():
    """Configure logging for the ETL process"""
    # TODO: Implement logging setup
    pass

def create_directory_structure():
    """Create necessary directories if they don't exist"""
    config = load_config()
    paths = config['paths']
    
    for path in paths.values():
        Path(path).mkdir(parents=True, exist_ok=True)