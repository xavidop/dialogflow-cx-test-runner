package cicd

import (
	"github.com/xavidop/dialogflow-cx-test-runner/internal/global"
	cxpkg "github.com/xavidop/dialogflow-cx-test-runner/pkg/cx"
)

func ExecutePipeline(envName string, locationID string, projectID string, agentName string) error {

	agentClient, err := cxpkg.CreateAgentClient(locationID)
	if err != nil {
		return err
	}
	defer agentClient.Close()

	agent, err := cxpkg.GetAgentIdByName(agentClient, agentName, projectID, locationID)
	if err != nil {
		return err
	}

	environmentClient, err := cxpkg.CreateEnvironmentsClient(locationID)
	if err != nil {
		return err
	}
	defer environmentClient.Close()

	env, err := cxpkg.GetEnvironmentIdByName(environmentClient, agent.GetName(), envName)
	if err != nil {
		return err
	}

	result, err := cxpkg.RunContinuousTest(environmentClient, env)
	if err != nil {
		return err
	}

	global.Log.Infof(result.GetResult().String())

	return nil
}
