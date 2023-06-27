import os
import json
import logging
import copy
from ruamel.yaml import YAML
from jsonpath_ng import parse

current_script_path = os.path.dirname(os.path.abspath(__file__))

variables_path = os.path.join(current_script_path, "../shared/variables.json")
config_template_path = os.path.join(current_script_path, "config_template.yaml")
config_path = os.path.join(current_script_path, "config.yaml")


def GenerateTargetList(network):
    network_name = network["name"]

    new_yaml_target = {
        "name": network_name,
        "rpcURL": network["url"],
    }

    yaml_data["TargetList"].append(new_yaml_target)


def GenerateTargetEventContracts(network, json_contracts, new_listening_data):
    aggregation_spotter_template = copy.deepcopy(
        yaml_data["ContractTemplates"]["AggregationTemplate"]
    )
    aggregation_spotter_template["address"] = json_contracts["aggregationSpotter"]

    new_listening_data["targetNetwork"] = {}
    new_listening_data["targetNetwork"]["rpcURL"] = network["url"]
    new_listening_data["targetNetwork"]["contracts"] = []
    new_listening_data["targetNetwork"]["contracts"].append(
        aggregation_spotter_template
    )

    return new_listening_data


def GenerateAggregationSpotterEvents(
    network, tent_network, json_contracts, ent_json_contracts, ent_spotters_json
):
    aggregation_spotter_template = copy.deepcopy(
        yaml_data["ContractTemplates"]["AggregationTemplate"]
    )

    new_listening_data = copy.deepcopy(
        yaml_data["ListeningDataTemplates"]["AggregationSpotter"]
    )

    new_listening_data["name"] = network["name"] + " AggregationSpotter"
    new_listening_data["rpcURL"] = network["url"]
    new_listening_data["contracts"] = []
    new_listening_data["contracts"].append(aggregation_spotter_template)
    new_listening_data["contracts"][0]["address"] = json_contracts["aggregationSpotter"]

    new_listening_data = GenerateTargetEventContracts(
        network, json_contracts, new_listening_data
    )

    return new_listening_data


def AddListningData(listeningData):
    yaml_data["ListeningData"].append(listeningData)


with open(variables_path, "r") as json_file:
    json_data = json.load(json_file)

yaml = YAML()
yaml.preserve_quotes = True
yaml.allow_duplicate_keys = True
with open(config_template_path, "r") as yaml_file:
    yaml_data = yaml.load(yaml_file)

oracleNetworks = {"ent"}
yaml_data["TargetList"] = []
yaml_data["ListeningData"] = []


json_networks = json_data["networks"]
tent_network = {}
for item in json_networks:
    if item.get("name") == "ent":
        tent_network = item
        break

for network in json_networks:
    network_name = network["name"]

    # json_contracts = json_data["deployedContractAddresses"][network_name]
    json_contracts = json_data["spotterAddresses"]["Addresses"][network_name]
    tent_json_contracts = json_data["deployedContractAddresses"]["ent"]
    ent_spotters_json = json_data["spotterAddresses"]["Addresses"]["ent"]
    # print("json_contracts:  ", json_contracts)
    logging.info("json_contracts: ", json_contracts)
    GenerateTargetList(network)

    # if network_name in oracleNetworks:
    AddListningData(
        GenerateAggregationSpotterEvents(
            network,
            tent_network,
            json_contracts,
            tent_json_contracts,
            ent_spotters_json,
        )
    )


with open(config_path, "w") as yaml_file:
    yaml.dump(yaml_data, yaml_file)
