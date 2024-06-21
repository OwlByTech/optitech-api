-- name: GetInstitutionClient :one
SELECT * FROM institution_client
WHERE client_id = $1 AND institution_id=$2 AND deleted_at IS NULL;

-- name: ListInstitutionClients :many
SELECT client.given_name,client.surname,client.email ,institution_client.client_id FROM institution_client
INNER JOIN client ON institution_client.client_id=client.client_id
WHERE  institution_client.institution_id=$1;


-- name: CreateInstitutionClient :copyfrom
INSERT INTO institution_client(client_id, institution_id, created_at)
VALUES ($1, $2, $3);

-- name: ExistsInstitutionClient :one
SELECT * FROM institution_client
WHERE client_id = $1 AND institution_id=$2 and deleted_at IS NOT NULL ;


-- name: RecoverInstitutionClient :exec
UPDATE institution_client
SET deleted_at = NULL,updated_at= $2
WHERE institution_id= $1 AND client_id= $3 ;


-- name: DeleteInstitutionByClient :exec
UPDATE institution_client
SET deleted_at = $2
WHERE institution_id= $1 AND client_id= $3 ;

-- name: DeleteInstitutionClientByInstitution :exec
UPDATE institution_client
SET deleted_at = $2
WHERE institution_id = $1;
