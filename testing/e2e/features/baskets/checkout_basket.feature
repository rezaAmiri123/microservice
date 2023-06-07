Feature: Checking out baskets
  As a basket with a basket I can check out

  Background:
    Given I am a registered user
    And I start a new basket
    And a store has the following items
      | Name              | Price |
      | Wizard w/ crystal | 9.99  |
    And I add the items
      | Name              | Quantity |
      | Wizard w/ crystal | 10       |
    And I authorize the payment

  Scenario: Checking out
    Given I check out the basket
    Then the basket is checked out
