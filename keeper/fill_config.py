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


def GenerateTargetList(network, network_contracts):
    network_name = network["name"]

    aggregation_spotter_template = copy.deepcopy(
        yaml_data["ContractTemplates"]["AggregationTemplate"]
    )

    new_yaml_target = {
        "name": network_name,
        "rpcURL": network["url"],
        "contracts": [],
    }

    new_yaml_target["contracts"].append(aggregation_spotter_template)
    new_yaml_target["contracts"][0]["address"] = network_contracts["aggregationSpotter"]

    yaml_data["TargetList"].append(new_yaml_target)


def GenerateTargetEventContracts(
    tent_network, ent_json_contracts, ent_spotters_json, new_listening_data
):
    aggregation_spotter_template = copy.deepcopy(
        yaml_data["ContractTemplates"]["AggregationTemplate"]
    )
    aggregation_spotter_template["address"] = ent_spotters_json["aggregationSpotter"]

    new_listening_data["targetNetwork"] = {}
    new_listening_data["targetNetwork"]["rpcURL"] = tent_network["url"]
    new_listening_data["targetNetwork"]["contracts"] = []
    new_listening_data["targetNetwork"]["contracts"].append(
        aggregation_spotter_template
    )

    return new_listening_data


def GenerateUSDTTransferEvents(
    network, tent_network, json_contracts, ent_json_contracts, ent_spotters_json
):
    usdt_template = copy.deepcopy(yaml_data["ContractTemplates"]["ETHUsdtTemplate"])

    new_listening_data = copy.deepcopy(yaml_data["ListeningDataTemplates"]["USDT"])

    new_listening_data["name"] = network_name + " USDT"
    new_listening_data["rpcURL"] = network["url"]
    new_listening_data["contracts"] = []
    new_listening_data["contracts"].append(usdt_template)
    new_listening_data["contracts"][0]["address"] = json_contracts["USDTContract"]

    new_listening_data = GenerateTargetEventContracts(
        tent_network, ent_json_contracts, ent_spotters_json, new_listening_data
    )

    return new_listening_data


# def GenerateBalanceSpotterEvents(network, ent_json_contracts, ent_spotters_json):
#     balance_spotter_template = copy.deepcopy(
#         yaml_data["ContractTemplates"]["BalanceSpotterTemplate"]
#     )

#     new_listening_data = copy.deepcopy(
#         yaml_data["ListeningDataTemplates"]["BalanceSpotter"]
#     )

#     new_listening_data["name"] = network_name + " BalanceSpotter"
#     new_listening_data["rpcURL"] = network["url"]
#     new_listening_data["contracts"] = []
#     new_listening_data["contracts"].append(balance_spotter_template)
#     new_listening_data["contracts"][0]["address"] = ent_spotters_json["balanceSpotter"]

#     return new_listening_data


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

    json_contracts = json_data["deployedContractAddresses"][network_name]
    ent_json_contracts = json_data["deployedContractAddresses"]["ent"]
    ent_spotters_json = json_data["spotterAddresses"]["Addresses"]["ent"]
    network_contracts = json_data["spotterAddresses"]["Addresses"][network_name]

    GenerateTargetList(network, network_contracts)

    # if network_name in oracleNetworks:
    #     AddListningData(
    #         GenerateBalanceSpotterEvents(network, ent_json_contracts, ent_spotters_json)
    #     )
    # else:
    if network_name not in oracleNetworks:
        AddListningData(
            GenerateUSDTTransferEvents(
                network,
                tent_network,
                json_contracts,
                ent_json_contracts,
                ent_spotters_json,
            )
        )


with open(config_path, "w") as yaml_file:
    yaml.dump(yaml_data, yaml_file)
