package types

// EspressoSchemaV1 Define the combined struct 
type EspressoSchemaV1 struct {
    EspressoTxResponseV1 EspressoTxResponseV1 `json:"espresso_tx_response_v_1"`
    StationId            string               `json:"station_id"`
}

type EspressoTxResponseV1 struct {
   Transaction Transaction `json:"transaction"`
   Hash        string      `json:"hash"`
   Index       int         `json:"index"`
   Proof       Proof       `json:"proof"`
   BlockHash   string      `json:"block_hash"`
   BlockHeight int         `json:"block_height"`
}

type Transaction struct {
   Namespace int    `json:"namespace"`
   Payload   string `json:"payload"`
}

type Proof struct {
   TxIndex                    string       `json:"tx_index"`
   PayloadNumTxs              string       `json:"payload_num_txs"`
   PayloadProofNumTxs         ProofDetails `json:"payload_proof_num_txs"`
   PayloadTxTableEntries      string       `json:"payload_tx_table_entries"`
   PayloadProofTxTableEntries ProofDetails `json:"payload_proof_tx_table_entries"`
   PayloadProofTx             ProofDetails `json:"payload_proof_tx"`
}

type ProofDetails struct {
   Proofs      string `json:"proofs"`
   PrefixBytes string `json:"prefix_bytes"`
   SuffixBytes string `json:"suffix_bytes"`
}