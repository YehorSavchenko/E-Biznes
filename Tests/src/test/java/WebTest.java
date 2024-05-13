import java.time.Duration;

import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.openqa.selenium.By;
import org.openqa.selenium.WebDriver;
import org.openqa.selenium.WebElement;
import org.openqa.selenium.chrome.ChromeDriver;
import org.openqa.selenium.support.ui.ExpectedConditions;
import org.openqa.selenium.support.ui.Select;
import org.openqa.selenium.support.ui.WebDriverWait;

import static org.junit.jupiter.api.Assertions.*;

public class WebTest {

    private WebDriver driver;

    @BeforeEach
    public void setup() {
        driver = new ChromeDriver();
        driver.manage().timeouts().implicitlyWait(Duration.ofMillis(500));
        driver.get("https://www.selenium.dev/selenium/web/web-form.html");
    }

    @Test
    public void testTextInput() {
        WebElement textBox = driver.findElement(By.name("my-text"));
        assertEquals("", textBox.getAttribute("value"));
        assertTrue(textBox.isEnabled(), "Text box should be enabled.");
        assertEquals("1px solid rgb(206, 212, 218)", textBox.getCssValue("border"));
        assertEquals("400", textBox.getCssValue("font-weight"));
        textBox.sendKeys("Test Input");
        assertEquals("Test Input", textBox.getAttribute("value"));
        textBox.clear();
        textBox.sendKeys("Another Test");
        assertEquals("Another Test", textBox.getAttribute("value"));
        assertEquals("", textBox.getAttribute("placeholder"));
        assertEquals("solid", textBox.getCssValue("border-style"));
    }

    @Test
    public void testTextInputWithSpecialCharacters() {
        WebElement textBox = driver.findElement(By.name("my-text"));
        String specialInput = "!@#$%^&*()_+{}|:\"<>?";
        textBox.sendKeys(specialInput);
        assertEquals(specialInput, textBox.getAttribute("value"));
        textBox.clear();
        assertEquals("", textBox.getAttribute("value"));
    }

    @Test
    public void testPasswordInput() {
        WebElement passwordBox = driver.findElement(By.name("my-password"));
        passwordBox.sendKeys("password123");
        assertEquals("password123", passwordBox.getAttribute("value"));
        assertEquals("password", passwordBox.getAttribute("type"));
        assertEquals(11, passwordBox.getAttribute("value").length());
    }

    @Test
    public void testTextarea() {
        WebElement textarea = driver.findElement(By.name("my-textarea"));
        textarea.sendKeys("Hello, this is a test.");
        assertEquals("Hello, this is a test.", textarea.getAttribute("value"));
    }

    @Test
    public void testTextareaWithNewlines() {
        WebElement textarea = driver.findElement(By.name("my-textarea"));
        String multilineText = "Line 1\nLine 2\nLine 3";
        textarea.sendKeys(multilineText);
        assertEquals(multilineText, textarea.getAttribute("value"));
    }

    @Test
    public void testElementStyles() {
        WebElement header = driver.findElement(By.tagName("h1"));
        assertEquals("start", header.getCssValue("text-align"), "Header text should be start");
        assertEquals("rgba(33, 37, 41, 1)", header.getCssValue("color"));
    }

    @Test
    public void testCheckbox() {
        WebElement checkbox = driver.findElement(By.id("my-check-1"));
        assertTrue(checkbox.isSelected());
        checkbox.click();
        assertFalse(checkbox.isSelected());
    }

    @Test
    public void testRadioButtons() {
        WebElement radio1 = driver.findElement(By.id("my-radio-1"));
        WebElement radio2 = driver.findElement(By.id("my-radio-2"));
        assertTrue(radio1.isSelected());
        radio1.click();
        assertFalse(radio2.isSelected());
        radio2.click();
        assertTrue(radio2.isSelected());
        assertFalse(radio1.isSelected());
    }

    @Test
    public void testDropdown() {
        Select dropdown = new Select(driver.findElement(By.name("my-select")));
        assertEquals(4, dropdown.getOptions().size());
        for (WebElement option : dropdown.getOptions()) {
            dropdown.selectByVisibleText(option.getText());
            assertEquals(option.getAttribute("value"), dropdown.getFirstSelectedOption().getAttribute("value"));
        }
        assertEquals("Three", dropdown.getFirstSelectedOption().getText());
    }

    @Test
    public void testRangeSlider() {
        WebElement range = driver.findElement(By.name("my-range"));
        range.sendKeys("5");
        assertEquals("5", range.getAttribute("value"));
    }

    @Test
    public void testDateInput() {
        WebElement dateInput = driver.findElement(By.name("my-date"));
        dateInput.sendKeys("2023-05-01");
        assertEquals("2023-05-01", dateInput.getAttribute("value"));
    }

    @Test
    public void testSubmitButton() {
        driver.findElement(By.cssSelector("button")).click();
        WebElement message = driver.findElement(By.id("message"));
        assertEquals("Received!", message.getText());
    }

    @Test
    public void testDisabledInput() {
        WebElement disabledInput = driver.findElement(By.name("my-disabled"));
        assertThrows(org.openqa.selenium.InvalidElementStateException.class,
                () -> disabledInput.sendKeys("Should fail"));
    }

    @Test
    public void testReadonlyInput() {
        WebElement readonlyInput = driver.findElement(By.name("my-readonly"));
        assertEquals("Readonly input", readonlyInput.getAttribute("value"));
    }

    @Test
    public void testDatalistInput() {
        WebElement datalistInput = driver.findElement(By.name("my-datalist"));
        datalistInput.sendKeys("New York");
        assertEquals("New York", datalistInput.getAttribute("value"));
    }

    @Test
    public void testColorPicker() {
        WebElement colorPicker = driver.findElement(By.name("my-colors"));
        colorPicker.sendKeys("#ff0000");
        assertEquals("#ff0000", colorPicker.getAttribute("value"));
    }

    @Test
    public void testDatePicker() {
        WebElement datePicker = driver.findElement(By.name("my-date"));
        datePicker.sendKeys("01/01/2023");
        assertEquals("01/01/2023", datePicker.getAttribute("value"));
    }

    @Test
    public void testUncheckedCheckbox() {
        WebElement checkbox2 = driver.findElement(By.id("my-check-2"));
        assertFalse(checkbox2.isSelected());
        checkbox2.click();
        assertTrue(checkbox2.isSelected());
    }

    @Test
    public void testDefaultRadioButton() {
        WebElement radio2 = driver.findElement(By.id("my-radio-2"));
        assertFalse(radio2.isSelected());
        radio2.click();
        assertTrue(radio2.isSelected());
    }

    @Test
    public void testHiddenInput() {
        WebElement hiddenInput = driver.findElement(By.name("my-hidden"));
        assertEquals("", hiddenInput.getAttribute("value"));
    }

    @Test
    public void testFormSubmission() {
        driver.findElement(By.id("my-text-id")).sendKeys("Testing Submission");
        driver.findElement(By.cssSelector("button[type='submit']")).click();
        assertNotEquals("Web form", driver.getTitle());
    }

    @Test
    public void testNavigationToIndexAndOverflowBody() {
        WebElement returnToIndexLink = driver.findElement(By.linkText("Return to index"));
        returnToIndexLink.click();

        WebDriverWait wait = new WebDriverWait(driver, Duration.ofSeconds(10));
        wait.until(ExpectedConditions.visibilityOfElementLocated(By.linkText("overflow-body.html")));

        assertEquals("Index of Available Pages", driver.getTitle());

        WebElement overflowLink = driver.findElement(By.linkText("overflow-body.html"));
        overflowLink.click();

        wait.until(ExpectedConditions.visibilityOfElementLocated(By.tagName("iframe")));
        WebElement iframe = driver.findElement(By.tagName("iframe"));
        assertTrue(iframe.isDisplayed());

        assertTrue(driver.getCurrentUrl().contains("overflow-body.html"));

        assertEquals("https://www.selenium.dev/selenium/web/xhtmlTest.html", iframe.getAttribute("src"));
    }

    @Test
    public void testContentOnOverflowBody() {
        driver.get("https://www.selenium.dev/selenium/web/overflow-body.html");

        assertEquals("The Visibility of Everyday Things", driver.getTitle());

        WebElement image = driver.findElement(By.cssSelector("img[alt='a nice beach']"));
        assertTrue(image.isDisplayed());

        assertNotEquals(0, image.getSize().height);
        assertNotEquals(0, image.getSize().width);

        WebElement iframe = driver.findElement(By.tagName("iframe"));
        driver.switchTo().frame(iframe);

        assertTrue(driver.findElement(By.tagName("body")).getText().contains("XHTML Might Be The Future"));

        WebElement h1 = driver.findElement(By.tagName("h1"));
        assertTrue(h1.isDisplayed());
        assertEquals("XHTML Might Be The Future", h1.getText());

        driver.switchTo().defaultContent();

        assertTrue(driver.findElements(By.tagName("button")).isEmpty());
    }

    @AfterEach
    public void teardown() {
        if (driver != null) {
            driver.quit();
        }
    }

}