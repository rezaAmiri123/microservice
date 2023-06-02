Feature: Starting baskets

  As a user I can start new shopping baskets

  Scenario: Create a basket
    Given I am a registered user
    When I start a new basket
    Then I expect the basket was started
