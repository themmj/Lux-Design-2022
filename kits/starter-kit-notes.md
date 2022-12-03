# Notes for future starter kits

There are a few gotchas one may run into when developing a starter kit. Most of them are relevant for statically
typed programming languages.

It is obviously recommended to use the existing kits as reference. But this list may still help to catch issues that
are not too obvious in the existing code.

- [1 Compiled kits](#1-compiled-kits)
- [2 main.py](#2-mainpy)
- [3 Input/Output](#3-inputoutput)
- [4 Configuration](#4-configuration)
- [5 Observation](#5-observation)
  - [5.1 Weather schedule](#51-weather-schedule)
  - [5.2 Board](#52-board)
  - [5.3 Factory occupancy](#53-factory-occupancy)
- [6 Actions](#6-actions)
  - [6.1 Bid action](#61-bid-action)
  - [6.2 Spawn action](#62-spawn-action)
  - [6.3 Factory actions](#63-factory-actions)
  - [6.4 Unit actions](#64-unit-actions)

## 1 Compiled kits

Kaggle servers currently run Ubuntu 18.04. Kits that require compilation to a native binary should support a build
using docker. That way everyone, no matter the OS, can submit their agent to Kaggle.
A native build should also be available for local testing.

## 2 main.py

To make the kit work on Kaggle the submission needs a `main.py` which creates a process from the binary built with
docker. Check the existing kits for reference of such a `main.py`.

## 3 Input/Output

Your kit receives input from `stdin` and communicates its output to `stdout`. The format used for communication is `JSON`.
Any form of logging (e.g. for debugging) can be directed to `stderr`.

Your agent will run as long as there is input available.
Every turn, your agent must provide output. (at least an empty `JSON` object)

## 4 Configuration

The configuration under `"info"."env_cfg"` is only available on the first step. (`"step"` = 0)

Weather information in this configuration under `"WEATHER"` always have an array of two integers called `"TIME_RANGE"`.
However, they *MAY* also contain a floating point number `"POWER_GAIN"` or `"POWER_CONSUMPTION"`. If they are not set
in the configuration, they should be 1 by default.

## 5 Observation

The observation under `"obs"` contains the observation for the current turn.

### 5.1 Weather schedule

The weather forecast under `"weather_schedule"` is only available on the first turn.

### 5.2 Board

In the first turn, the `"board"` will contain all members as described in the other documentation. In the following
turns, the board input will only contain `"factories_per_team"`, `"lichen"`, `"lichen_strains"` and `"rubble"`.

Information like the lichen, ore, ice etc. should be represented as a 2D array of integers, indexed `[x][y]`.
You'll receive this data as such on the first turn. However, for all following turns the remaining attributes will be a
`JSON` object resembling a sparse matrix for updates. Every attribute is a string containing the x and y coordinate
separated by a comma. The corresponding value is the updated value. E.g.

```json
"rubble" {
    "1,5": 1,
    "3,8": 10
}
```

Also note that the `"valid_spawns_mask"` is also a 2D array, but contains *booleans* instead of integers.

### 5.3 Factory occupancy

If you decide to introduce a factory occupancy map for convenience, keep in mind to *reset* and then repopulate it based
on the current list of factories. This will take into account that a factory may be destroyed.

## 6 Actions

There are a few different action types, and they differ how lux expects them to be communicated.

Bid and Spawn actions can be returned as their `JSON` representation respectively.

Factory and Unit actions are returned in the form of a `JSON` object which maps the units' ids to their action queues.
E.g.

```json
{
    "unit_0": [ ... ],
    "unit_5": [ ... ]
}
```

### 6.1 Bid action

A simple object with:
- `faction`: string
- `bid`: integer

### 6.2 Spawn action

A simple object with:
- `spawn`: array of 2 integers (x and y)
- `metal`: integer
- `water`: integer

### 6.3 Factory actions

A factory action is a single integer which represents the desired action:
- `0`: build one light unit
- `1`: build one heavy unit
- `2`: water the lichen

### 6.4 Unit actions

Each unit action is an array of 5 integers.

0. Type of the action
1. Direction
2. distance or Resource (for Transfer and Pickup actions)
3. amount
4. repeat (how many times it's supposed to be repeated)

Type:
- `0`: Move
- `1`: Transfer
- `2`: Pickup
- `3`: Dig
- `4`: Self destruct
- `5`: Recharge

Direction:
- `0`: Center
- `1`: Up
- `2`: Right
- `3`: Down
- `4`: Left

Resource:
- `0`: Ice
- `1`: Ore
- `2`: Water
- `3`: Metal
- `4`: Power

