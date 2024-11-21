-- Variants table for storing gnomAD data
CREATE TABLE IF NOT EXISTS variants (
    id SERIAL PRIMARY KEY,
    chromosome VARCHAR(10),
    position INTEGER,
    reference_allele TEXT,
    alternate_allele TEXT,
    gene_symbol VARCHAR(100),
    variant_id VARCHAR(100) UNIQUE,
    allele_frequency FLOAT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Population-specific metrics
CREATE TABLE IF NOT EXISTS population_metrics (
    id SERIAL PRIMARY KEY,
    variant_id VARCHAR(100) REFERENCES variants(variant_id),
    population VARCHAR(50),
    allele_count INTEGER,
    allele_frequency FLOAT,
    homozygote_count INTEGER
);

-- Create indexes for efficient querying
CREATE INDEX idx_variants_gene ON variants(gene_symbol);
CREATE INDEX idx_variants_position ON variants(chromosome, position);
CREATE INDEX idx_population_variant ON population_metrics(variant_id);