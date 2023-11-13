# hlf-qscc-parser
Hyperledger Fabric QSCC Parser in Golang: Decoding the transaction, block output from QSCC to a readable format.

## How to use

1. Call GetTransactionByID

```bash
peer chaincode query -o localhost:7050 -C mychannel -n qscc -c '{"function":"GetTransactionByID","Args":["mychannel", "<txId>"]}' --tls --cafile "${PWD}"/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem --hex > tx.txt
```

2. Unmarshall the transaction information to readable format

```bash
cat tx.txt | go run .
```

3. Example output

```json
{
	"ValidationCode": 0,
	"TransactionEnvelope": {
		"Payload": {
			"Header": {
				"ChannelHeader": {
					"Type": 3,
					"Version": 0,
					"Timestamp": {
						"seconds": 1698204767,
						"nanos": 871425295
					},
					"ChannelId": "mychannel",
					"TxId": "928ea3337e5edba4ce967b434c7fe48ae7deda3873286f782a6f438959d6d391",
					"Epoch": 0,
					"Extension": "EgkSB3ByaXZhdGU=",
					"TlsCertHash": null
				},
				"SignatureHeader": {
					"Creator": {
						"Mspid": "Org1MSP",
						"IdBytes": {
							"Signature": "MEUCIQDntTiKYGVReW5V/K6SaqZwQWKwSK0IE6zxDY23nwI+ygIgWpbMcCmt9U8e2WWMtiZ2EPujjkbVmcFXc9Afh8WZ4UY=",
							"SignatureAlgorithm": "ECDSA-SHA256",
							"PublicKeyAlgorithm": "ECDSA",
							"PublicKey": {
								"Curve": {},
								"X": 104150099026742245382696428945300298186287598685286875582450075765946126100215,
								"Y": 9892313427572036350611029736725131491917076334091913992234785027541066528607
							},
							"Version": 3,
							"SerialNumber": 707524332766909212973881292688936508588524226041,
							"Issuer": "CN=ca.org1.example.com,O=org1.example.com,L=Durham,ST=North Carolina,C=US",
							"Subject": "CN=org1admin,OU=admin,O=Hyperledger,ST=North Carolina,C=US",
							"NotBefore": "2023-10-25T03:26:00Z",
							"NotAfter": "2024-10-24T03:31:00Z",
							"KeyUsage": 1,
							"Extensions": [
								{
									"Id": "2.5.29.15",
									"Critical": true,
									"Value": "AwIHgA=="
								},
								{
									"Id": "2.5.29.19",
									"Critical": true,
									"Value": "MAA="
								},
								{
									"Id": "2.5.29.14",
									"Critical": false,
									"Value": "BBR7WaLlfphfkkwt+h5s0H/Roqy4rg=="
								},
								{
									"Id": "2.5.29.35",
									"Critical": false,
									"Value": "MBaAFKa/+BJnma6716QYizz9puOQv55i"
								},
								{
									"Id": "2.5.29.17",
									"Critical": false,
									"Value": "MAWCA0t1bg=="
								},
								{
									"Id": "1.2.3.4.5.6.7.8.1",
									"Critical": false,
									"Value": "eyJhdHRycyI6eyJoZi5BZmZpbGlhdGlvbiI6IiIsImhmLkVucm9sbG1lbnRJRCI6Im9yZzFhZG1pbiIsImhmLlR5cGUiOiJhZG1pbiJ9fQ=="
								}
							],
							"BasicConstraintsValid": true,
							"IsCA": false,
							"MaxPathLen": -1,
							"MaxPathLenZero": false,
							"SubjectKeyId": "e1mi5X6YX5JMLfoebNB/0aKsuK4=",
							"AuthorityKeyId": "pr/4EmeZrrvXpBiLPP2m45C/nmI="
						}
					},
					"Nonce": "DLJRqh371z2Ra16aP9Qfm4+/qNGwRBv/"
				}
			},
			"Data": {
				"Actions": [
					{
						"Header": {
							"Creator": {
								"Mspid": "Org1MSP",
								"IdBytes": {
									"Signature": "MEUCIQDntTiKYGVReW5V/K6SaqZwQWKwSK0IE6zxDY23nwI+ygIgWpbMcCmt9U8e2WWMtiZ2EPujjkbVmcFXc9Afh8WZ4UY=",
									"SignatureAlgorithm": "ECDSA-SHA256",
									"PublicKeyAlgorithm": "ECDSA",
									"PublicKey": {
										"Curve": {},
										"X": 104150099026742245382696428945300298186287598685286875582450075765946126100215,
										"Y": 9892313427572036350611029736725131491917076334091913992234785027541066528607
									},
									"Version": 3,
									"SerialNumber": 707524332766909212973881292688936508588524226041,
									"Issuer": "CN=ca.org1.example.com,O=org1.example.com,L=Durham,ST=North Carolina,C=US",
									"Subject": "CN=org1admin,OU=admin,O=Hyperledger,ST=North Carolina,C=US",
									"NotBefore": "2023-10-25T03:26:00Z",
									"NotAfter": "2024-10-24T03:31:00Z",
									"KeyUsage": 1,
									"Extensions": [
										{
											"Id": "2.5.29.15",
											"Critical": true,
											"Value": "AwIHgA=="
										},
										{
											"Id": "2.5.29.19",
											"Critical": true,
											"Value": "MAA="
										},
										{
											"Id": "2.5.29.14",
											"Critical": false,
											"Value": "BBR7WaLlfphfkkwt+h5s0H/Roqy4rg=="
										},
										{
											"Id": "2.5.29.35",
											"Critical": false,
											"Value": "MBaAFKa/+BJnma6716QYizz9puOQv55i"
										},
										{
											"Id": "2.5.29.17",
											"Critical": false,
											"Value": "MAWCA0t1bg=="
										},
										{
											"Id": "1.2.3.4.5.6.7.8.1",
											"Critical": false,
											"Value": "eyJhdHRycyI6eyJoZi5BZmZpbGlhdGlvbiI6IiIsImhmLkVucm9sbG1lbnRJRCI6Im9yZzFhZG1pbiIsImhmLlR5cGUiOiJhZG1pbiJ9fQ=="
										}
									],
									"BasicConstraintsValid": true,
									"IsCA": false,
									"MaxPathLen": -1,
									"MaxPathLenZero": false,
									"SubjectKeyId": "e1mi5X6YX5JMLfoebNB/0aKsuK4=",
									"AuthorityKeyId": "pr/4EmeZrrvXpBiLPP2m45C/nmI="
								}
							},
							"Nonce": "DLJRqh371z2Ra16aP9Qfm4+/qNGwRBv/"
						},
						"Payload": {
							"ChaincodeProposalPayload": {
								"Input": {
									"ChaincodeSpec": {
										"Type": "GOLANG",
										"ChaincodeId": {
											"Name": "private",
											"Version": "",
											"Path": ""
										},
										"Input": {
											"Args": {
												"Args": [
													"CreateAsset"
												]
											},
											"Decorations": null,
											"IsInit": false
										},
										"Timeout": 0
									}
								},
								"TransientMap": null
							},
							"Action": {
								"ProposalResponsePayload": {
									"ProposalHash": "gOo/0tunCRTYp/kUBfN8MGIApviW40uAHJsCGvv1f9E=",
									"Extension": {
										"Results": {
											"DataModel": "KV",
											"NsRwset": [
												{
													"Namespace": "_lifecycle",
													"Rwset": {
														"Reads": [
															{
																"Key": "namespaces/fields/private/Collections",
																"Version": {
																	"BlockNum": 5,
																	"TxNum": 0
																}
															},
															{
																"Key": "namespaces/fields/private/EndorsementInfo",
																"Version": {
																	"BlockNum": 5,
																	"TxNum": 0
																}
															},
															{
																"Key": "namespaces/fields/private/Sequence",
																"Version": {
																	"BlockNum": 5,
																	"TxNum": 0
																}
															},
															{
																"Key": "namespaces/fields/private/ValidationInfo",
																"Version": {
																	"BlockNum": 5,
																	"TxNum": 0
																}
															},
															{
																"Key": "namespaces/metadata/private",
																"Version": {
																	"BlockNum": 5,
																	"TxNum": 0
																}
															}
														],
														"RangeQueriesInfo": [],
														"Writes": [],
														"MetadataWrites": []
													},
													"CollectionHashedRwset": []
												},
												{
													"Namespace": "private",
													"Rwset": {
														"Reads": [],
														"RangeQueriesInfo": [],
														"Writes": [],
														"MetadataWrites": []
													},
													"CollectionHashedRwset": [
														{
															"CollectionName": "Org1MSPPrivateCollection",
															"HashedRwset": {
																"HashedReads": [],
																"HashedWrites": [
																	{
																		"KeyHash": "jj3S6p/z2nCGKlJiH3wdyBwrGEy4hqMko/Qw7BHv0/I=",
																		"IsDelete": false,
																		"ValueHash": "/DN+C/CcRaHLMGXHJet89mIUZ5RRGtUDVjr0sdDdP0c=",
																		"IsPurge": false
																	}
																],
																"MetadataWrites": []
															},
															"PvtRwsetHash": "Q6t2gpAi3QdM7DH8emgU6dDpJkibQCnuQQLQd8dUejE="
														},
														{
															"CollectionName": "assetCollection",
															"HashedRwset": {
																"HashedReads": [
																	{
																		"KeyHash": "jj3S6p/z2nCGKlJiH3wdyBwrGEy4hqMko/Qw7BHv0/I=",
																		"Version": {
																			"BlockNum": 0,
																			"TxNum": 0
																		}
																	}
																],
																"HashedWrites": [
																	{
																		"KeyHash": "jj3S6p/z2nCGKlJiH3wdyBwrGEy4hqMko/Qw7BHv0/I=",
																		"IsDelete": false,
																		"ValueHash": "I0ezfLxd/Pro7OpnZNGwe+9VSQLDt1lEc5Jlc2Uv8oY=",
																		"IsPurge": false
																	}
																],
																"MetadataWrites": []
															},
															"PvtRwsetHash": "AriKup/RZ+j1sRm6LQh6h34B2fVDH/cwwYOZEC0U1Po="
														}
													]
												}
											]
										},
										"Events": {
											"ChaincodeId": "",
											"TxId": "",
											"EventName": "",
											"Payload": null
										},
										"Response": {
											"Status": 200,
											"Message": "",
											"Payload": null
										},
										"ChaincodeId": {
											"Name": "private",
											"Version": "1.0",
											"Path": ""
										}
									}
								},
								"Endorsements": [
									{
										"Endorser": {
											"Mspid": "Org1MSP",
											"IdBytes": {
												"Signature": "MEQCIC+T5fhW+6p1ekt1dzETp4lNDXfiHFRUKyBcndlid8/hAiAYBPuPo1dZweQey4gzXz6XDF2Jw2HEYQirY9GDXQhrZw==",
												"SignatureAlgorithm": "ECDSA-SHA256",
												"PublicKeyAlgorithm": "ECDSA",
												"PublicKey": {
													"Curve": {},
													"X": 30881680888579190587092773701199638097513710429716051986852593777332631772509,
													"Y": 71515445683894570937045372659875161678822092804882224248507605195395442576759
												},
												"Version": 3,
												"SerialNumber": 721821521795934294922685582002796989684170750982,
												"Issuer": "CN=ca.org1.example.com,O=org1.example.com,L=Durham,ST=North Carolina,C=US",
												"Subject": "CN=peer0,OU=peer,O=Hyperledger,ST=North Carolina,C=US",
												"NotBefore": "2023-10-25T03:26:00Z",
												"NotAfter": "2024-10-24T03:31:00Z",
												"KeyUsage": 1,
												"Extensions": [
													{
														"Id": "2.5.29.15",
														"Critical": true,
														"Value": "AwIHgA=="
													},
													{
														"Id": "2.5.29.19",
														"Critical": true,
														"Value": "MAA="
													},
													{
														"Id": "2.5.29.14",
														"Critical": false,
														"Value": "BBRvyD3ae2Sxnzxs+efzcWBQqXI3dA=="
													},
													{
														"Id": "2.5.29.35",
														"Critical": false,
														"Value": "MBaAFKa/+BJnma6716QYizz9puOQv55i"
													},
													{
														"Id": "2.5.29.17",
														"Critical": false,
														"Value": "MBiCFnBlZXIwLm9yZzEuZXhhbXBsZS5jb20="
													},
													{
														"Id": "1.2.3.4.5.6.7.8.1",
														"Critical": false,
														"Value": "eyJhdHRycyI6eyJoZi5BZmZpbGlhdGlvbiI6IiIsImhmLkVucm9sbG1lbnRJRCI6InBlZXIwIiwiaGYuVHlwZSI6InBlZXIifX0="
													}
												],
												"BasicConstraintsValid": true,
												"IsCA": false,
												"MaxPathLen": -1,
												"MaxPathLenZero": false,
												"SubjectKeyId": "b8g92ntksZ88bPnn83FgUKlyN3Q=",
												"AuthorityKeyId": "pr/4EmeZrrvXpBiLPP2m45C/nmI="
											}
										},
										"Signature": "MEMCH35mcVANNU3FLiM8JfdAWn3aSZOyo3QPNTD2b97+nFcCIEPosGOAtn5Wti9P87q5dahr4/Z0c0j+h25ftN1tpvqz"
									}
								]
							}
						}
					}
				]
			}
		},
		"Signature": "MEUCIQDqT4dBQU3wSrW5c+Bg53GqE8fpZcSsdrGZt3sJOz8kNQIgKJNWMkd9dcD1aNG4kJxEXEg1klatJhllXZqk0Kpb0xQ="
	}
}
```
