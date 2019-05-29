package eth

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	log "github.com/golang/glog"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/joincivil/go-common/pkg/generated/contract"
)

// RetryCivilTCRContract is a CivilTCRContract that includes functions to
// make contract calls that retry.
type RetryCivilTCRContract struct {
	*contract.CivilTCRContract
}

// ChallengesWithRetry is a version of Challenges that has a retry mechanism
// Will retry on empty values for challenges or on error
func (r *RetryCivilTCRContract) ChallengesWithRetry(opts *bind.CallOpts, challengeID *big.Int,
	maxAttempts int, baseWaitMs int) (struct {
	RewardPool  *big.Int
	Challenger  common.Address
	Resolved    bool
	Stake       *big.Int
	TotalTokens *big.Int
}, error) {

	var challengeRes struct {
		RewardPool  *big.Int
		Challenger  common.Address
		Resolved    bool
		Stake       *big.Int
		TotalTokens *big.Int
	}

	attempt := 1
	var err error
	for {
		retry := false

		challengeRes, err = r.Challenges(opts, challengeID)

		// If error, try again
		// If challenge data looks empty, try again
		if err != nil {
			log.Errorf(
				"challenge error %v, sleep/attempt again, waiting %v ms...",
				err,
				baseWaitMs*attempt,
			)
			retry = true

		} else if challengeRes.Stake.Int64() == 0 && challengeRes.RewardPool.Int64() == 0 {
			log.Infof(
				"challenge does not look ready, sleep/attempt again, waiting %v ms...",
				baseWaitMs*attempt,
			)
			retry = true
		}

		if retry {
			// Take a break and see if the data propagates
			time.Sleep(time.Duration(baseWaitMs) * time.Duration(attempt) * time.Millisecond)
			if attempt > maxAttempts {
				if err != nil {
					return challengeRes, fmt.Errorf("exceeded max attempts to get challenges: err: %v", err)
				}
				return challengeRes, errors.New("exceeded max attempts to get challenges")
			}
			attempt++
			continue
		}

		break
	}

	return challengeRes, nil
}

// ChallengeRequestAppealExpiriesWithRetry is a version of ChallengeRequestAppealExpiriesWithRetry
// that has a retry mechanism
// Will retry on empty values for challenges or on error
func (r *RetryCivilTCRContract) ChallengeRequestAppealExpiriesWithRetry(opts *bind.CallOpts,
	challengeID *big.Int, maxAttempts int, baseWaitMs int) (*big.Int, error) {
	attempt := 1
	var err error
	var expiry *big.Int
	for {
		retry := false

		expiry, err = r.ChallengeRequestAppealExpiries(opts, challengeID)

		// If error, try again
		// If expiry looks empty, try again
		if err != nil {
			log.Errorf(
				"challenge appeal expiry error %v, sleep/attempt again, waiting %v ms...",
				err,
				baseWaitMs*attempt,
			)
			retry = true

		} else if expiry.Int64() == 0 {
			log.Infof(
				"challenge appeal expiry does not look ready, sleep/attempt again, waiting %v ms...",
				baseWaitMs*attempt,
			)
			retry = true
		}

		if retry {
			// Take a break and see if the data propagates
			time.Sleep(time.Duration(baseWaitMs) * time.Duration(attempt) * time.Millisecond)
			if attempt > maxAttempts {
				if err != nil {
					return nil, fmt.Errorf("exceeded max attempts to get challenge appeal expiry: err: %v", err)
				}
				return nil, errors.New("exceeded max attempts to get challenge appeal expiry")
			}
			attempt++
			continue
		}

		break
	}

	return expiry, nil

}
