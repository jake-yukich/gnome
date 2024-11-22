// Package variant handles all variant-related functionality for gnome.
//
// This file defines the Variant and PopulationMetric structs that:
// 1. Match the database schema defined in db/migrations/
// 2. Will be used by both the API handlers and ETL pipeline
// 3. Provide JSON serialization for API responses
//
// Related files:
// - db/migrations/001_variants.up.sql: Defines the matching tables
// - db/schema.go: Creates these tables in the database
package variant

import "time"

type Variant struct {
	ID              int       `json:"id"`
	Chromosome      string    `json:"chromosome"`
	Position        int       `json:"position"`
	ReferenceAllele string    `json:"reference_allele"`
	AlternateAllele string    `json:"alternate_allele"`
	GeneSymbol      string    `json:"gene_symbol"`
	VariantID       string    `json:"variant_id"`
	AlleleFrequency float64   `json:"allele_frequency"`
	CreatedAt       time.Time `json:"created_at"`
}

type PopulationMetric struct {
	ID              int     `json:"id"`
	VariantID       string  `json:"variant_id"`
	Population      string  `json:"population"`
	AlleleCount     int     `json:"allele_count"`
	AlleleFrequency float64 `json:"allele_frequency"`
	HomozygoteCount int     `json:"homozygote_count"`
}
