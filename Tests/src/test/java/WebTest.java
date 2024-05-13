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
        textBox.sendKeys("Test Input");
        assertEquals("Test Input", textBox.getAttribute("value"));
    }

    @Test
    public void testPasswordInput() {
        WebElement passwordBox = driver.findElement(By.name("my-password"));
        passwordBox.sendKeys("password123");
        assertEquals("password123", passwordBox.getAttribute("value"));
    }

    @Test
    public void testTextarea() {
        WebElement textarea = driver.findElement(By.name("my-textarea"));
        textarea.sendKeys("Hello, this is a test.");
        assertEquals("Hello, this is a test.", textarea.getAttribute("value"));
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
        dropdown.selectByVisibleText("Two");
        assertEquals("2", dropdown.getFirstSelectedOption().getAttribute("value"));
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

        new WebDriverWait(driver, Duration.ofSeconds(10)).until(ExpectedConditions.visibilityOfElementLocated(By.linkText("overflow-body.html")));

        WebElement overflowLink = driver.findElement(By.linkText("overflow-body.html"));
        overflowLink.click();

        new WebDriverWait(driver, Duration.ofSeconds(10))
                .until(ExpectedConditions.visibilityOfElementLocated(By.tagName("iframe")));

        assertTrue(driver.findElement(By.tagName("iframe")).isDisplayed(), "Iframe should be visible on 'overflow-body.html'");
    }

    @Test
    public void testContentOnOverflowBody() {
        driver.get("https://www.selenium.dev/selenium/web/overflow-body.html");

        WebElement image = driver.findElement(By.cssSelector("img[alt='a nice beach']"));
        assertTrue(image.isDisplayed(), "Image with alt text 'a nice beach' should be displayed");

        WebElement iframe = driver.findElement(By.tagName("iframe"));
        driver.switchTo().frame(iframe);
        assertTrue(driver.findElement(By.tagName("body")).getText().contains("XHTML Might Be The Future"),
                "Text within the iframe should match expected content");
    }

    @AfterEach
    public void teardown() {
        if (driver != null) {
            driver.quit();
        }
    }

}