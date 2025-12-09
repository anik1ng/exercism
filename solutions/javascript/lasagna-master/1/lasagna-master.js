/// <reference path="./global.d.ts" />
// @ts-check

/**
 * Implement the functions needed to solve the exercise here.
 * Do not forget to export them so they are available for the
 * tests. Here an example of the syntax as reminder:
 *
 * export function yourFunction(...) {
 *   ...
 * }
 */

export function cookingStatus(remainingTime) {
  if (remainingTime === undefined) {
    return 'You forgot to set the timer.'
  }

  if (remainingTime === 0) {
    return 'Lasagna is done.'
  }

  return 'Not done, please wait.'
}

export function preparationTime(layers, averageTime = 2) {
  return layers.length * averageTime
}

export function quantities(layers) {
  const result = {
    'noodles': 0,
    'sauce': 0
  }

  for (let layer of layers) {
    switch (layer) {
      case 'noodles':
        result.noodles += 50
        break
      case 'sauce':
        result.sauce += 0.2
        break
    }
  }

  return result
}

export function addSecretIngredient(friendsList, myList) {
  myList.push(friendsList.at(-1))
}

export function scaleRecipe(recipe, quantities) {
  const result = {}
  
  for (let ingridient in recipe) {
    result[ingridient] = recipe[ingridient] * (quantities / 2)
  }
  
  return result
}

