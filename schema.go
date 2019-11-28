package client

import (
	"fmt"
)

// JobFragment Reusable job fragment
const JobFragment = `
			id
			status
			request {
				type
				callbackURL
				properties{
					name
					value
				}
				parameters {
					idPhoto
					userPhoto
					firstName
					lastName
					dob
				}
			}
			result {
				success
				type
				country
				state

				id
				firstName
				lastName
				middleName
				birthDate
				expireDate

				confidences {
					id
					backId
					selfie
					idMatch
					faceMatch
				}
			}
			errors {
				type
				message
				suggestion
			}
			submitted
`

var (
	// JobsQuery GraphQL query for accessing jobs
	JobsQuery = fmt.Sprintf(`
	query jobs(
    $id: ID
    $ids: [ID]
    $type: String
    $token: String
    $status: String
    $to: String
    $from: String
    $withPhotos: Boolean
    $sortOrder: String
    $sortBy: String
    $page: Int
    $pageSize: Int
  ) {
    jobs(
      withPhotos: $withPhotos
      id: $id
      ids: $ids
      status: $status
      type: $type
      token: $token
      to: $to
      from: $from
      sortOrder: $sortOrder
      sortBy: $sortBy
      page: $page
      pageSize: $pageSize
    ) {
      total
      totalPages
      pageSize
      page
      items {
        %s
      }
    }
  }
	`, JobFragment)

	// UpdateSecretClientKeyMutation GraphQL mutation for submitting jobs
	UpdateSecretClientKeyMutation = `
		mutation updateSecretClientKey($secretClientKey: String) {
			updateSecretClientKey(secretClientKey: $secretClientKey) {
				secretClientKey
			}
		}
	`
	// SubmitMutation GraphQL mutation for submitting jobs
	SubmitMutation = fmt.Sprintf(`
		mutation submitJob(
			$type: String!
			$callbackURL: String
			$properties: [JobPropertyParam]
			$params: JobParams
		) {
			submitJob(
				type: $type
				callbackURL: $callbackURL
				properties: $properties
				params: $params
			) {
				%s
			}
		}
	`, JobFragment)

	// RemoveMutation GraphQL mutation for removing jobs
	RemoveMutation = fmt.Sprintf(`
	  mutation removeJob(
			$id: ID!
			) {
				removeJob(
					id: $id
				) {
					%s
				}
		}
		`, JobFragment)
)
